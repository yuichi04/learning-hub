# データ構造

## 概要

データ構造は、データを効率的に格納、操作、アクセスするための方法です。適切なデータ構造を選択することで、アルゴリズムの性能を大幅に向上させることができます。このノートでは、一般的なデータ構造とその基本操作について説明します。

## 配列 (Array)

配列は、固定サイズの連続したメモリ領域にデータを格納するデータ構造です。

### 基本操作

-   **アクセス**: 配列の任意の要素に直接アクセスできます。
-   **挿入**: 任意の位置に要素を挿入できますが、効率は位置によります。
-   **削除**: 任意の位置から要素を削除できますが、効率は位置によります。

### サンプルコード

```javascript
let arr = [1, 2, 3, 4, 5];

// アクセス
console.log(arr[2]); // 出力: 3

// 挿入
arr.splice(2, 0, 6);
console.log(arr); // 出力: [1, 2, 6, 3, 4, 5]

// 削除
arr.splice(3, 1);
console.log(arr); // 出力: [1, 2, 6, 4, 5]
```

## リスト (List)

リストは、要素が線形に並んだデータ構造であり、各要素が次の要素へのポインタを持つことで実現されます。

### 単方向リスト (Singly Linked List)

単方向リストは、各要素が次の要素へのポインタを持つリストです。

### 基本操作

-   **挿入**: リストの先頭や末尾、または任意の位置に要素を挿入できます。
-   **削除**: リストの任意の位置から要素を削除できます。

### サンプルコード

```javascript
class Node {
    constructor(value) {
        this.value = value;
        this.next = null;
    }
}

class SinglyLinkedList {
    constructor() {
        this.head = null;
    }

    insert(value) {
        let newNode = new Node(value);
        if (!this.head) {
            this.head = newNode;
        } else {
            let current = this.head;
            while (current.next) {
                current = current.next;
            }
            current.next = newNode;
        }
    }

    delete(value) {
        if (!this.head) return;

        if (this.head.value === value) {
            this.head = this.head.next;
            return;
        }

        let current = this.head;
        while (current.next && current.next.value !== value) {
            current = current.next;
        }

        if (current.next) {
            current.next = current.next.next;
        }
    }

    display() {
        let current = this.head;
        while (current) {
            console.log(current.value);
            current = current.next;
        }
    }
}

let list = new SinglyLinkedList();
list.insert(1);
list.insert(2);
list.insert(3);
list.display();
// 出力:
// 1
// 2
// 3

list.delete(2);
list.display();
// 出力:
// 1
// 3
```

## スタック (Stack)

スタックは、LIFO (Last In, First Out) 方式で要素を管理するデータ構造です。

### 基本操作

-   **プッシュ**: スタックのトップに要素を追加します。
-   **ポップ**: スタックのトップから要素を削除します。
-   **ピーク**: スタックのトップの要素を参照しますが、削除はしません。

### サンプルコード

```javascript
class Stack {
    constructor() {
        this.items = [];
    }

    push(element) {
        this.items.push(element);
    }

    pop() {
        if (this.items.length === 0) return "Underflow";
        return this.items.pop();
    }

    peek() {
        return this.items[this.items.length - 1];
    }

    isEmpty() {
        return this.items.length === 0;
    }

    display() {
        console.log(this.items.toString());
    }
}

let stack = new Stack();
stack.push(10);
stack.push(20);
stack.push(30);
stack.display(); // 出力: 10,20,30

console.log(stack.pop()); // 出力: 30
stack.display(); // 出力: 10,20
```

## キュー (Queue)

キューは、FIFO (First In, First Out) 方式で要素を管理するデータ構造です。

### 基本操作

-   **エンキュー**: キューの末尾に要素を追加します。
-   **デキュー**: キューの先頭から要素を削除します。
-   **フロント**: キューの先頭の要素を参照しますが、削除はしません。

### サンプルコード

```javascript
class Queue {
    constructor() {
        this.items = [];
    }

    enqueue(element) {
        this.items.push(element);
    }

    dequeue() {
        if (this.items.length === 0) return "Underflow";
        return this.items.shift();
    }

    front() {
        if (this.items.length === 0) return "No elements in Queue";
        return this.items[0];
    }

    isEmpty() {
        return this.items.length === 0;
    }

    display() {
        console.log(this.items.toString());
    }
}

let queue = new Queue();
queue.enqueue(10);
queue.enqueue(20);
queue.enqueue(30);
queue.display(); // 出力: 10,20,30

console.log(queue.dequeue()); // 出力: 10
queue.display(); // 出力: 20,30
```

## ツリー (Tree)

ツリーは、階層的にデータを管理するデータ構造で、各ノードが複数の子ノードを持つことができます。

### 二分木 (Binary Tree)

二分木は、各ノードが最大で 2 つの子ノードを持つツリーです。

### サンプルコード

```javascript
class TreeNode {
    constructor(value) {
        this.value = value;
        this.left = null;
        this.right = null;
    }
}

class BinaryTree {
    constructor() {
        this.root = null;
    }

    insert(value) {
        let newNode = new TreeNode(value);
        if (this.root === null) {
            this.root = newNode;
        } else {
            this.insertNode(this.root, newNode);
        }
    }

    insertNode(node, newNode) {
        if (newNode.value < node.value) {
            if (node.left === null) node.left = newNode;
            else this.insertNode(node.left, newNode);
        } else {
            if (node.right === null) node.right = newNode;
            else this.insertNode(node.right, newNode);
        }
    }

    inOrder(node) {
        if (node !== null) {
            this.inOrder(node.left);
            console.log(node.value);
            this.inOrder(node.right);
        }
    }

    getRootNode() {
        return this.root;
    }
}

let tree = new BinaryTree();
tree.insert(10);
tree.insert(20);
tree.insert(5);
tree.insert(15);

let root = tree.getRootNode();
tree.inOrder(root);
// 出力:
// 5
// 10
// 15
// 20
```

## グラフ (Graph)

グラフは、頂点（ノード）とそれらを結ぶ辺（エッジ）から構成されるデータ構造です。

### サンプルコード

```javascript
class Graph {
    constructor() {
        this.adjList = new Map();
    }

    addVertex(v) {
        this.adjList.set(v, []);
    }

    addEdge(v, w) {
        this.adjList.get(v).push(w);
        this.adjList.get(w).push(v);
    }

    printGraph() {
        let keys = this.adjList.keys();

        for (let i of keys) {
            let values = this.adjList.get(i);
            let conc = "";
            for (let j of values) {
                conc += j + " ";
            }
            console.log(i + " -> " + conc);
        }
    }
}

let graph = new Graph();
graph.addVertex("A");
graph.addVertex("B");
graph.addVertex("C");
graph.addEdge("A", "B");
graph.addEdge("A", "C");
graph.addEdge("B", "C");
graph.printGraph();
// 出力:
// A -> B C
// B -> A C
// C -> A B
```

## 結論

データ構造は、効率的なデータ管理と操作のための重要な概念です。各データ構造の特性と基本操作を理解することで、適切なデータ構造を選択し、アルゴリズムの効率を最大化することができます。
