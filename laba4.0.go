package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Функция для вычисления абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Генерация случайного массива
func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(2000) - 1000 // Генерация чисел от -1000 до 1000
	}
	return arr
}

// Создание частично отсортированного массива
func generatePartiallySortedArray(arr []int, sortedPercentage float64) []int {
	size := len(arr)
	sortedSize := int(float64(size) * sortedPercentage / 100)
	partiallySortedArr := make([]int, size)
	copy(partiallySortedArr, arr)

	// Сортируем первые sortedSize элементов по абсолютным значениям
	for i := 1; i < sortedSize; i++ {
		key := partiallySortedArr[i]
		j := i - 1
		for j >= 0 && abs(partiallySortedArr[j]) > abs(key) {
			partiallySortedArr[j+1] = partiallySortedArr[j]
			j--
		}
		partiallySortedArr[j+1] = key
	}
	return partiallySortedArr
}

// Создание массива, отсортированного в обратном порядке
func generateReversedArray(arr []int) []int {
	reversedArr := make([]int, len(arr))
	copy(reversedArr, arr)
	for i, j := 0, len(reversedArr)-1; i < j; i, j = i+1, j-1 {
		reversedArr[i], reversedArr[j] = reversedArr[j], reversedArr[i]
	}
	return reversedArr
}

// Сортировка с помощью прямого обмена (пузырьковая сортировка)
func bubbleSort(arr []int) (time.Duration, int, int) {
	start := time.Now()
	n := len(arr)
	comparisons := 0
	swaps := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if abs(arr[j]) > abs(arr[j+1]) { // Сравниваем абсолютные значения
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
			}
		}
	}
	return time.Since(start), comparisons, swaps
}

// Сортировка с помощью прямого включения (сортировка вставками)
func insertionSort(arr []int) (time.Duration, int, int) {
	start := time.Now()
	n := len(arr)
	comparisons := 0
	swaps := 0
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && abs(arr[j]) > abs(key) { // Сравниваем абсолютные значения
			comparisons++
			arr[j+1] = arr[j]
			j--
			swaps++
		}
		arr[j+1] = key
	}
	return time.Since(start), comparisons, swaps
}

// Сортировка с помощью прямого выбора (сортировка выбором)
func selectionSort(arr []int) (time.Duration, int, int) {
	start := time.Now()
	n := len(arr)
	comparisons := 0
	swaps := 0
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			comparisons++
			if abs(arr[j]) < abs(arr[minIndex]) { // Сравниваем абсолютные значения
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
			swaps++
		}
	}
	return time.Since(start), comparisons, swaps
}

// Быстрая сортировка (QuickSort)
func quickSort(arr []int, low, high int) (time.Duration, int, int) {
	start := time.Now()
	comparisons, swaps := quickSortHelper(arr, low, high)
	return time.Since(start), comparisons, swaps
}

// Вспомогательная функция для быстрой сортировки
func quickSortHelper(arr []int, low, high int) (int, int) {
	comparisons := 0
	swaps := 0
	if low < high {
		pivot, cmp, swp := partition(arr, low, high)
		comparisons += cmp
		swaps += swp
		cmp1, swp1 := quickSortHelper(arr, low, pivot-1)
		cmp2, swp2 := quickSortHelper(arr, pivot+1, high)
		comparisons += cmp1 + cmp2
		swaps += swp1 + swp2
	}
	return comparisons, swaps
}

// Разделение массива для быстрой сортировки
func partition(arr []int, low, high int) (int, int, int) {
	pivot := arr[high]
	i := low - 1
	comparisons := 0
	swaps := 0
	for j := low; j < high; j++ {
		comparisons++
		if abs(arr[j]) < abs(pivot) { // Сравниваем абсолютные значения
			i++
			arr[i], arr[j] = arr[j], arr[i]
			swaps++
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	swaps++
	return i + 1, comparisons, swaps
}

// Функция для пользовательского ввода целого числа
func getUserInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n') // Читаем ввод пользователя
		if err != nil {
			fmt.Println("Ошибка чтения ввода. Попробуйте снова.")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Ошибка: ввод не может быть пустым. Введите целое число больше 0.")
			continue
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите только целое число больше 0.")
			continue
		}
		if num <= 0 {
			fmt.Println("Ошибка: размер массива должен быть больше 0.")
			continue
		}
		return num
	}
}

// Тестирование алгоритмов сортировки
func testSortingAlgorithms(arr []int) {
	// Алгоритмы сортировки
	algorithms := []struct {
		name string
		sort func([]int) (time.Duration, int, int)
	}{
		{"Сортировка с помощью прямого обмена", bubbleSort},
		{"Сортировка с помощью прямого включения", insertionSort},
		{"Сортировка с помощью прямого выбора", selectionSort},
		{"Быстрая сортировка", func(arr []int) (time.Duration, int, int) {
			return quickSort(arr, 0, len(arr)-1)
		}},
	}

	// Типы массивов
	types := []struct {
		name string
		data []int
	}{
		{"", arr},
		{"Частично отсортированный массив (25%)", generatePartiallySortedArray(arr, 25)},
		{"Частично отсортированный массив (50%)", generatePartiallySortedArray(arr, 50)},
		{"Частично отсортированный массив (75%)", generatePartiallySortedArray(arr, 75)},
		{"Отсортированный в обратном порядке массив", generateReversedArray(arr)},
	}

	// Тестирование
	for _, t := range types {
		fmt.Printf("\n--- %s ---\n", t.name)
		for _, algo := range algorithms {
			arrCopy := make([]int, len(t.data))
			copy(arrCopy, t.data)
			duration, comparisons, swaps := algo.sort(arrCopy)
			fmt.Printf("\n%s\n", algo.name)
			fmt.Printf("Отсортированный массив: %v\n", arrCopy)
			fmt.Printf("Время выполнения: %v, Сравнения: %d, Перестановки: %d\n",
				duration, comparisons, swaps)
		}
	}
}

func main() {
	// Пользовательский ввод размера массива
	size := getUserInput("Введите размер массива: ")

	// Генерация случайного массива
	randomArray := generateRandomArray(size)
	fmt.Println("\nСгенерированный массив:")
	fmt.Println(randomArray)

	// Тестирование алгоритмов сортировки
	testSortingAlgorithms(randomArray)
}
