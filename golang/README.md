# Go事前課題

この事前課題は、以下のことを目的としています。

- Go言語の言語機能の習得

## 課題: A Tour of Go

[A Tour of Go](https://tour.golang.org/) は、Go言語の基本的な言語機能を体験できる資料です。

A Tour of Goを一通り読み、途中で出題されるExerciseに回答してください。Exerciseは、本リポジトリの`exercise-n`ディレクトリ中に同じ内容のファイルがあるので、これを編集して回答してください。

#### Basics

- [Excercise: Loops and Functions](https://tour.golang.org/flowcontrol/8) ... [`exercise-1`](./exercise-1)
- [Excercise: Slices](https://tour.golang.org/moretypes/18) ... [`exercise-2`](./exercise-2)
- [Excercise: Maps](https://tour.golang.org/moretypes/23) ... [`exercise-3`](./exercise-3)
- [Excercise: Fibonacci closure](https://tour.golang.org/moretypes/26) ... [`exercise-4`](./exercise-4)

#### Methods and interfaces

- [Excercise: Stringers](https://tour.golang.org/methods/18) ... [`exercise-5`](./exercise-5)
- [Excercise: Errors](https://tour.golang.org/methods/20) ... [`exercise-6`](./exercise-6)
- [Excercise: Readers](https://tour.golang.org/methods/22) ... [`exercise-7`](./exercise-7)
- [Excercise: rot13Reader](https://tour.golang.org/methods/23) ... [`exercise-8`](./exercise-8)
- [Excercise: Images](https://tour.golang.org/methods/25) ... [`exercise-9`](./exercise-9)

#### Concurrency

[Concurrency](https://tour.golang.org/concurrency/1) はオプション課題とします。任意で取り組んでください。

- [Excercise: Equivalent Binary Trees](https://tour.golang.org/concurrency/7)
- [Excersise: Web Crawler](https://tour.golang.org/concurrency/10)

### 取り組み方

リポジトリのルートで、以下のように個別のexerciseを実行できます。

```console
$ make go-run-exercise-1
```

期待通りに動作しているかどうかは、リポジトリのルートで、以下のようにテストを実行することで確認できます。

```console
$ make go-test
```

## 参考資料

- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
    - Go言語のイディオムが紹介されている
