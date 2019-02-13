package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setupParams(fuelСonsumption, pricePerOneLiter float64, distance int) float64 {

	var distanceInt = float64(distance)
	consumptionOfLitersPerOneKilometer := fuelСonsumption / 100
	totalCosts := consumptionOfLitersPerOneKilometer * pricePerOneLiter * distanceInt
	totalLiters := distanceInt / 100 / fuelСonsumption

	fmt.Println("\nРезультат:")
	t := fmt.Sprintf("Количество литров на %.d километров, л: %.1f", distance, totalLiters)
	fmt.Println(t)

	return totalCosts
}

const s = "Неверный воод! Введите число > 0"

func main() {
	consolescanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Добро пожаловать в программу \"Расчет стоимости поездки")

	for {

		distance := inputDestantion(consolescanner)
		for distance <= 0 {
			fmt.Println(s)
			distance = inputDestantion(consolescanner)
		}

		fuelСonsumption := fuelConsumption(consolescanner)
		for fuelСonsumption <= 0 {
			fmt.Println(s)
			fuelСonsumption = fuelConsumption(consolescanner)

		}
		pricePerOneLiter := price(consolescanner)
		for pricePerOneLiter <= 0 {
			fmt.Println(s)
			pricePerOneLiter = price(consolescanner)
		}

		if err := consolescanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		result := setupParams(fuelСonsumption, pricePerOneLiter, distance)
		tc := fmt.Sprintf("Стоимость топлива на %.d километров, грн: %.2f ", distance, result)
		fmt.Println(tc)

		for {
			fmt.Print("\nПродолжить? (y/n или да/нет): ")
			consolescanner.Scan()
			nextStep := strings.ToLower(consolescanner.Text())

			if nextStep == "y" || nextStep == "да" {
				break
			} else if nextStep == "n" || nextStep == "нет" {
				fmt.Println("До свидания!")
				return
			} else {
				fmt.Println("Неверный ввод! ")
			}
		}
	}
}

func price(consolescanner *bufio.Scanner) float64 {
	fmt.Print("Введите стоимость 1-ого литра в гривнах : ")
	consolescanner.Scan()
	pricePerOneLiter, _ := strconv.ParseFloat(consolescanner.Text(), 64)
	return pricePerOneLiter
}

func fuelConsumption(consolescanner *bufio.Scanner) float64 {
	fmt.Print("Введите расход топлива на 100 км : ")
	consolescanner.Scan()
	fuelСonsumption, _ := strconv.ParseFloat(consolescanner.Text(), 64)
	return fuelСonsumption
}

func inputDestantion(consolescanner *bufio.Scanner) int {
	fmt.Print("\nВведите дистанцию в км : ")
	consolescanner.Scan()
	distance, _ := strconv.Atoi(consolescanner.Text())
	return distance
}
