package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Пример: Аукцион
Аукцион состоит из нескольких участников, которые делают постоянно увеличивающиеся ставки по лоту
Он может закончиться в следующих случаях:
1. Дана максимальная ставка
2. Сделано определнное кол-во ставок
3. Прошло определенное время
*/

type Lot struct {
	sync.Mutex
	CurrentBid int
	PlayerID   int
	MaxPrice   int
	MaxBids    int
	currentCnt int
}

func (b *Lot) GetCurrentBid() int {
	b.Lock()
	defer b.Unlock()
	return b.CurrentBid
}

func (b *Lot) SetNewBid(newBid PlayerBid) bool {
	b.Lock()
	defer b.Unlock()
	if newBid.Bid > b.CurrentBid {
		fmt.Printf("new bid: %+v\n", newBid)

		b.CurrentBid = newBid.Bid
		b.PlayerID = newBid.PlayerID
		b.currentCnt++
		//finish auction if called maximum price
		if b.MaxBids <= b.currentCnt {
			println("finish by count")
			return true
		}
		//finish auction if the maximum bets were
		if b.MaxPrice <= newBid.Bid {
			println("finish by bid")
			return true
		}
	}
	return false
}

type PlayerBid struct {
	Bid      int
	PlayerID int
}

const (
	AvgSleep   = 20000
	AvgBidStep = 100
)

func makeBid(ctx context.Context, playerID int, lot *Lot, bids chan PlayerBid) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		myBid := lot.GetCurrentBid() + rand.Intn(AvgBidStep)

		select {
		case <-ctx.Done():
			return
		case bids <- PlayerBid{PlayerID: playerID, Bid: myBid}:
		default:
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(AvgSleep)))
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	lot := &Lot{CurrentBid: 100, MaxBids: 100, MaxPrice: 10000}
	bids := make(chan PlayerBid)
	//create context which will signal in certain time
	ctx, finish := context.WithTimeout(context.Background(), time.Second*20)

	//run 5 players who will create bets
	for i := 0; i < 5; i++ {
		go makeBid(ctx, i, lot, bids)
	}

LOOP:
	for {
		select {
		case bid := <-bids:
			if lot.SetNewBid(bid) {
				finish()
				break LOOP
			}
		case <-ctx.Done():
			fmt.Println("Done", ctx.Err())
			break LOOP
		}
	}
	fmt.Println("Auction finished by player", lot.PlayerID, "with price", lot.CurrentBid)
}
