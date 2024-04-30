// package main

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	s := make(chan os.Signal, 1)
// 	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

// 	filename := "test.txt"
// 	errCh := make(chan error, 1)
// 	go func() {
// 		errCh <- writeFile(ctx, filename)
// 	}()

// 	select {
// 	case error := <-errCh:
// 		if error != nil {
// 			fmt.Printf("Error: %v\n", error)
// 		} else {
// 			fmt.Println("File created successfully and data written.")
// 		}
// 	case num := <-s:
// 		fmt.Printf("Signal received: %v\n", num)
// 		cancel()
// 		<-errCh
// 		fmt.Println("The program has been cancelled.")
// 	}
// }

// func writeFile(ctx context.Context, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.WriteString(file, "This is the data written to the file.\n")
// 	if err != nil {
// 		return err
// 	}

// 	select {
// 	case <-time.After(3 * time.Second):
// 		return nil
// 	case <-ctx.Done():
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		default:
// 			return fmt.Errorf("error writing file")
// 		}
// 	}
// }




package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func Read(ctx context.Context, filename string) ([]byte, error) {
	f, error := os.Open(filename)
	if error != nil {
		return nil, error
	}
	defer f.Close()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	data := make([]byte, 1024)
	x, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	return data[:x], nil
}

func Write(ctx context.Context, filename string, data []byte) error {
	f2, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f2.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, err = f2.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testFile := "test.txt"
	test2File := "test2.txt"

	Data, err := Read(ctx, testFile)
	if err != nil {
		log.Fatal(err)
	}

	err = Write(ctx, test2File, Data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully written to file.")
}
