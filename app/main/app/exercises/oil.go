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
	oneLitersForKilometers := 100 / fuelСonsumption
	totalLiters := distanceInt / oneLitersForKilometers

	//q := fmt.Sprintf("Quantity liters for 1 km: %.2f%s", consumptionOfLitersPerOneKilometer, "l")
	//q := fmt.Sprintf("Количество литров на 1 км: %.2f%s", consumptionOfLitersPerOneKilometer, "л")
	//t := fmt.Sprintf("Total liters for %.d kilometers: %.1f%s", distance, totalLiters, "l")
	fmt.Println("\nРезультат:")
	t := fmt.Sprintf("Количество литров на %.d километров: %.1f%s", distance, totalLiters, "л")
	//fmt.Println(q)
	fmt.Println(t)

	return totalCosts
}

/*
	consolescanner.Scan()

	switch input := consolescanner.Text(); input {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
*/

func main() {
	consolescanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Welcome to the \"Calculating the cost of the trip\" program")
	fmt.Println("Добро пожаловать в программу \"Расчет стоимости поездки\"\n")
	// by default, bufio.Scanner scans newline-separated lines

	for {
		fmt.Print("Введите дистанцию в км : ")
		consolescanner.Scan()
		inputDistance := consolescanner.Text()

		fmt.Print("Введите расход топлива на 100 км : ")
		consolescanner.Scan()
		inputfuelСonsumption := consolescanner.Text()

		fmt.Print("Введите стоимость 1-ого литра в гривнах : ")
		consolescanner.Scan()
		inputpricePerOneLiter := consolescanner.Text()

		if err := consolescanner.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		distance, _ := strconv.Atoi(inputDistance)
		fuelСonsumption, _ := strconv.ParseFloat(inputfuelСonsumption, 64)
		pricePerOneLiter, _ := strconv.ParseFloat(inputpricePerOneLiter, 64)
		result := setupParams(fuelСonsumption, pricePerOneLiter, distance)
		//tc := fmt.Sprintf("Total costs for %.d kilometers: %.2fuah ", distance, result)
		tc := fmt.Sprintf("Стоимость топлива на %.d километров: %.2fгрн. ", distance, result)
		fmt.Println(tc)

		for {
			fmt.Print("\nПродолжить? (y/n или да/нет): ")
			consolescanner.Scan()
			nextStep := consolescanner.Text()
			nextStep = strings.ToLower(nextStep)

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
