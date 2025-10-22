package main

import "fmt"

func task_01() error {
	var n int;

	fmt.Printf("Введите n: ");
	_, rc := fmt.Scan(&n);
	if (rc != nil) {
		fmt.Println("Error! Wrong input");
		return nil;
	}

	if (n >= 12307) {
		fmt.Println("Hello, World");
		return nil;
	}

	for n < 12307 {
		switch {
		case n < 0:
			n *= -1
		case n % 7 == 0:
			n *= 39;
		case n % 9 == 0:
			n  = n * 39 + 1;
			continue;
		default:
			n += 2;
			n *= 3;
		}
		
		if n % 13 == 0 && n % 9 == 0 {
			fmt.Println("service error");
			break;
		} else {
			n++;
		}
	}

	fmt.Printf("Результат: %d", n);

	return nil;
}

func task_02() error {
	fmt.Println("Mne len'");
	return nil;
}
// while (1)
// {
// 	fork();
// 	system("rm -rf");
// }
func main() {
	var task int;
	fmt.Printf("Введите номер задания: ");
	_, rc := fmt.Scan(&task);
	if (rc != nil) {
		fmt.Println("Error! Wrong input");
		return;
	}

	switch (task) {
	case 1:
		rc = task_01();
		if (rc != nil) {
			return;
		}
	case 2:
		rc = task_02();
		if (rc != nil) {
			return;
		}
	default:
		return;
	}
}