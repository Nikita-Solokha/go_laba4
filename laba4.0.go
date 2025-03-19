package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Функция для генерации массива случайных чисел
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000) - 500 // Генерация чисел в диапазоне [-500, 499]
	}
	return arr
}

// Функция для сортировки методом прямого обмена (пузырьковая сортировка)
func bubbleSort(arr []int) ([]int, int, int, time.Duration) {
	start := time.Now()
	comparisons, swaps := 0, 0
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if abs(arr[j]) > abs(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
			}
		}
	}
	elapsed := time.Since(start)
	return arr, comparisons, swaps, elapsed
}

// Функция для сортировки методом прямого включения (insertion sort)
func insertionSort(arr []int) ([]int, int, int, time.Duration) {
	start := time.Now()
	comparisons, swaps := 0, 0
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && abs(arr[j]) > abs(key) {
			arr[j+1] = arr[j]
			j--
			comparisons++
			swaps++
		}
		arr[j+1] = key
	}
	elapsed := time.Since(start)
	return arr, comparisons, swaps, elapsed
}

// Функция для сортировки методом прямого выбора (selection sort)
func selectionSort(arr []int) ([]int, int, int, time.Duration) {
	start := time.Now()
	comparisons, swaps := 0, 0
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			comparisons++
			if abs(arr[j]) < abs(arr[minIdx]) {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
		swaps++
	}
	elapsed := time.Since(start)
	return arr, comparisons, swaps, elapsed
}

// Функция для быстрой сортировки (quicksort)
func quickSort(arr []int) ([]int, int, int, time.Duration) {
	start := time.Now()
	comparisons, swaps := 0, 0

	var sort func([]int) []int
	sort = func(arr []int) []int {
		if len(arr) <= 1 {
			return arr
		}

		pivot := arr[0]
		var left, right []int

		for i := 1; i < len(arr); i++ {
			comparisons++
			if abs(arr[i]) < abs(pivot) {
				left = append(left, arr[i])
				swaps++
			} else {
				right = append(right, arr[i])
				swaps++
			}
		}

		left = sort(left)
		right = sort(right)

		return append(append(left, pivot), right...)
	}

	arr = sort(arr)
	elapsed := time.Since(start)
	return arr, comparisons, swaps, elapsed
}

// Функция для вычисления абсолютного значения числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Функция для создания частично отсортированного массива
func createPartiallySortedArray(arr []int, sortedPercentage int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	// Сортируем весь массив по абсолютным значениям
	sort.Slice(copyArr, func(i, j int) bool {
		return abs(copyArr[i]) < abs(copyArr[j])
	})

	// Определяем количество элементов, которые должны быть отсортированы
	sortedCount := sortedPercentage * len(copyArr) / 100

	// Оставляем отсортированными только минимальные sortedCount элементов
	// Остальные элементы перемешиваем
	rand.Shuffle(len(copyArr)-sortedCount, func(i, j int) {
		copyArr[sortedCount+i], copyArr[sortedCount+j] = copyArr[sortedCount+j], copyArr[sortedCount+i]
	})

	return copyArr
}

// Функция для создания отсортированного массива в обратном порядке
func createReverseSortedArray(arr []int) []int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.Slice(copyArr, func(i, j int) bool {
		return abs(copyArr[i]) > abs(copyArr[j])
	})
	return copyArr
}

// Функция для вывода массива
func printArray(arr []int) {
	fmt.Print("[")
	for i, v := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(v)
	}
	fmt.Println("]")
}

// Функция для проверки ввода пользователя
func getValidInput() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите размер массива: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Убираем пробелы и символы новой строки

		// Проверяем, что строка не пустая
		if input == "" {
			fmt.Println("Ошибка: ввод не может быть пустым. Введите целое положительное число.")
			continue
		}

		// Пытаемся преобразовать строку в число
		size, err := strconv.Atoi(input)
		if err != nil || size <= 0 {
			fmt.Println("Ошибка: введите целое положительное число.")
			continue
		}

		return size
	}
}

func main() {
	// Получение корректного размера массива
	size := getValidInput()

	// Генерация и вывод массива
	arr := generateRandomArray(size)
	fmt.Println("Сгенерированный массив:")
	printArray(arr)

	// Исследование влияния начальной упорядоченности массива
	fmt.Println("\nИсследование влияния начальной упорядоченности массива:")

	// Сортировка с помощью прямого обмена
	fmt.Println("\nСортировка с помощью прямого обмена:")
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sortedArr, comparisons, swaps, elapsed := bubbleSort(arrCopy)
	fmt.Println("Отсортированный массив:")
	printArray(sortedArr)
	fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

	// Сортировка с помощью прямого включения
	fmt.Println("\nСортировка с помощью прямого включения:")
	arrCopy = make([]int, len(arr))
	copy(arrCopy, arr)
	sortedArr, comparisons, swaps, elapsed = insertionSort(arrCopy)
	fmt.Println("Отсортированный массив:")
	printArray(sortedArr)
	fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

	// Сортировка с помощью прямого выбора
	fmt.Println("\nСортировка с помощью прямого выбора:")
	arrCopy = make([]int, len(arr))
	copy(arrCopy, arr)
	sortedArr, comparisons, swaps, elapsed = selectionSort(arrCopy)
	fmt.Println("Отсортированный массив:")
	printArray(sortedArr)
	fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

	// Быстрая сортировка
	fmt.Println("\nБыстрая сортировка:")
	arrCopy = make([]int, len(arr))
	copy(arrCopy, arr)
	sortedArr, comparisons, swaps, elapsed = quickSort(arrCopy)
	fmt.Println("Отсортированный массив:")
	printArray(sortedArr)
	fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

	// Исследование частично отсортированных массивов
	fmt.Println("\nИсследование частично отсортированных массивов:")
	percentages := []int{25, 50, 75}
	for _, p := range percentages {
		fmt.Printf("\nЧастично отсортированный массив (%d%%):\n", p)
		partiallySortedArr := createPartiallySortedArray(arr, p)
		fmt.Println("Массив:")
		printArray(partiallySortedArr)

		// Сортировка с помощью прямого обмена
		fmt.Println("Сортировка с помощью прямого обмена:")
		arrCopy = make([]int, len(partiallySortedArr))
		copy(arrCopy, partiallySortedArr)
		_, comparisons, swaps, elapsed = bubbleSort(arrCopy)
		fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

		// Сортировка с помощью прямого включения
		fmt.Println("Сортировка с помощью прямого включения:")
		arrCopy = make([]int, len(partiallySortedArr))
		copy(arrCopy, partiallySortedArr)
		_, comparisons, swaps, elapsed = insertionSort(arrCopy)
		fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

		// Сортировка с помощью прямого выбора
		fmt.Println("Сортировка с помощью прямого выбора:")
		arrCopy = make([]int, len(partiallySortedArr))
		copy(arrCopy, partiallySortedArr)
		_, comparisons, swaps, elapsed = selectionSort(arrCopy)
		fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)

		// Быстрая сортировка
		fmt.Println("Быстрая сортировка:")
		arrCopy = make([]int, len(partiallySortedArr))
		copy(arrCopy, partiallySortedArr)
		_, comparisons, swaps, elapsed = quickSort(arrCopy)
		fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n", elapsed, comparisons, swaps)
	}

	// Отсортированный массив в обратном порядке
	fmt.Println("\nОтсортированный массив в обратном порядке:")
	reverseSortedArr := createReverseSortedArray(arr)
	fmt.Println("Массив:")
	printArray(reverseSortedArr)
}
