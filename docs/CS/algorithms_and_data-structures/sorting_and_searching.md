# ソートと探索

## 概要

ソートと探索は、コンピュータサイエンスにおける基本的かつ重要な操作です。ソートはデータの順序を整えることで、探索はデータの中から特定の要素を見つけることを指します。このノートでは、一般的なソートと探索のアルゴリズムについて説明します。

## ソートアルゴリズム

### バブルソート (Bubble Sort)

バブルソートは、隣接する要素を比較し、必要に応じて交換することでリストをソートするシンプルなアルゴリズムです。

```javascript
function bubbleSort(arr) {
    let n = arr.length;
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
            }
        }
    }
    return arr;
}
```

### 選択ソート (Selection Sort)

選択ソートは、未ソート部分から最小（または最大）の要素を選んで、それをソート済み部分の末尾に追加するアルゴリズムです。

```javascript
function selectionSort(arr) {
    let n = arr.length;
    for (let i = 0; i < n - 1; i++) {
        let minIndex = i;
        for (let j = i + 1; j < n; j++) {
            if (arr[j] < arr[minIndex]) {
                minIndex = j;
            }
        }
        [arr[i], arr[minIndex]] = [arr[minIndex], arr[i]];
    }
    return arr;
}
```

### 挿入ソート (Insertion Sort)

挿入ソートは、未ソート部分から要素を一つずつ取り出し、それをソート済み部分の適切な位置に挿入するアルゴリズムです。

```javascript
function insertionSort(arr) {
    let n = arr.length;
    for (let i = 1; i < n; i++) {
        let key = arr[i];
        let j = i - 1;
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
    return arr;
}
```

### クイックソート (Quick Sort)

クイックソートは、基準となるピボットを選び、それより小さい要素と大きい要素に分けて再帰的にソートするアルゴリズムです。

```javascript
function quickSort(arr) {
    if (arr.length <= 1) {
        return arr;
    }
    let pivot = arr[Math.floor(arr.length / 2)];
    let left = [];
    let right = [];
    for (let i = 0; i < arr.length; i++) {
        if (i !== Math.floor(arr.length / 2)) {
            if (arr[i] < pivot) {
                left.push(arr[i]);
            } else {
                right.push(arr[i]);
            }
        }
    }
    return quickSort(left).concat([pivot], quickSort(right));
}
```

## 探索アルゴリズム

### 線形探索 (Linear Search)

線形探索は、リスト内の要素を最初から順番に比較して目的の要素を探すアルゴリズムです。

```javascript
function linearSearch(arr, target) {
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] === target) {
            return i;
        }
    }
    return -1;
}
```

### 二分探索 (Binary Search)

二分探索は、ソートされたリストを対象に、効率的に要素を探すアルゴリズムです。

```javascript
function binarySearch(arr, target) {
    let left = 0;
    let right = arr.length - 1;
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        if (arr[mid] === target) {
            return mid;
        } else if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}
```

## 結論

ソートと探索は、データ処理において非常に重要な操作です。これらのアルゴリズムを理解し、適切に選択・実装することで、効率的なプログラムを作成することができます。ここで紹介したアルゴリズムは、基本的なものであり、さらに高度なアルゴリズムを学ぶための基礎となります。
