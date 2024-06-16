# 再帰と分割統治法

## 概要

再帰と分割統治法は、複雑な問題を効率的に解決するための重要な手法です。再帰は、関数が自分自身を呼び出す手法であり、分割統治法は問題を小さな部分に分割して解決する手法です。このノートでは、これらの概念と具体的な例について説明します。

## 再帰 (Recursion)

### 再帰の基本概念

再帰は、関数が自分自身を呼び出すプロセスです。再帰的なアルゴリズムは、ベースケース（終了条件）と再帰ステップ（自己呼び出し）で構成されます。

### 再帰の例

#### 階乗 (Factorial)

階乗は、ある数のすべての正整数の積です。n の階乗は n \* (n-1)の階乗です。

```javascript
function factorial(n) {
    if (n === 0) {
        return 1; // ベースケース
    }
    return n * factorial(n - 1); // 再帰ステップ
}
```

#### フィボナッチ数列 (Fibonacci Sequence)

フィボナッチ数列は、各項が前の二つの項の和になる数列です。

```javascript
function fibonacci(n) {
    if (n <= 1) {
        return n; // ベースケース
    }
    return fibonacci(n - 1) + fibonacci(n - 2); // 再帰ステップ
}
```

## 分割統治法 (Divide and Conquer)

### 分割統治法の基本概念

分割統治法は、問題をより小さな部分に分割し、それぞれを再帰的に解決し、その結果を統合する手法です。一般的に、以下の 3 つのステップで構成されます：

1. **分割 (Divide)**: 問題を小さな部分に分割する。
2. **統治 (Conquer)**: 各部分を再帰的に解決する。
3. **統合 (Combine)**: 部分問題の解決結果を統合して元の問題の解を得る。

### 分割統治法の例

#### マージソート (Merge Sort)

マージソートは、リストを再帰的に分割し、マージしてソートするアルゴリズムです。

```javascript
function mergeSort(arr) {
    if (arr.length <= 1) {
        return arr; // ベースケース
    }

    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));
    const right = mergeSort(arr.slice(mid));

    return merge(left, right);
}

function merge(left, right) {
    let result = [];
    let i = 0;
    let j = 0;

    while (i < left.length && j < right.length) {
        if (left[i] < right[j]) {
            result.push(left[i]);
            i++;
        } else {
            result.push(right[j]);
            j++;
        }
    }

    return result.concat(left.slice(i)).concat(right.slice(j));
}
```

#### クイックソート (Quick Sort)

クイックソートは、ピボットを選択し、リストを再帰的にソートするアルゴリズムです。

```javascript
function quickSort(arr) {
    if (arr.length <= 1) {
        return arr; // ベースケース
    }

    const pivot = arr[Math.floor(arr.length / 2)];
    const left = [];
    const right = [];
    const equal = [];

    for (let num of arr) {
        if (num < pivot) {
            left.push(num);
        } else if (num > pivot) {
            right.push(num);
        } else {
            equal.push(num);
        }
    }

    return quickSort(left).concat(equal).concat(quickSort(right));
}
```

## 結論

再帰と分割統治法は、複雑な問題を効率的に解決するための強力な手法です。これらの手法を理解し適用することで、より効率的で洗練されたアルゴリズムを設計することができます。
