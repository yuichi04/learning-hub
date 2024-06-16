# グラフアルゴリズム

## 概要

グラフアルゴリズムは、グラフ構造を操作し、分析するためのアルゴリズムです。グラフは、頂点（ノード）とそれらを結ぶ辺（エッジ）から構成され、様々な現実世界の問題をモデル化するのに使用されます。このノートでは、代表的なグラフアルゴリズムとその基本操作について説明します。

## グラフの表現方法

グラフは主に 2 つの方法で表現されます：

-   **隣接リスト**: 各頂点に対して、接続している頂点のリストを保持します。
-   **隣接行列**: 頂点のペアごとに、辺の有無を示す行列を使用します。

### サンプルコード：グラフの隣接リスト表現

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
        this.adjList.get(w).push(v); // 無向グラフの場合
    }

    printGraph() {
        let keys = this.adjList.keys();
        for (let key of keys) {
            let values = this.adjList.get(key);
            let conc = "";
            for (let value of values) {
                conc += value + " ";
            }
            console.log(key + " -> " + conc);
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

## 深さ優先探索 (DFS: Depth First Search)

DFS は、グラフの頂点を深く探索していくアルゴリズムです。スタックを使って実装することが一般的です。

### サンプルコード：深さ優先探索

```javascript
class Graph {
    // ... 前述のGraphクラスのコード ...

    dfs(startingNode) {
        let visited = new Set();
        this._dfsHelper(startingNode, visited);
    }

    _dfsHelper(node, visited) {
        visited.add(node);
        console.log(node);

        let neighbors = this.adjList.get(node);
        for (let neighbor of neighbors) {
            if (!visited.has(neighbor)) {
                this._dfsHelper(neighbor, visited);
            }
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
graph.dfs("A");
// 出力:
// A
// B
// C
```

## 幅優先探索 (BFS: Breadth First Search)

BFS は、グラフの頂点を広く探索していくアルゴリズムです。キューを使って実装することが一般的です。

### サンプルコード：幅優先探索

```javascript
class Graph {
    // ... 前述のGraphクラスのコード ...

    bfs(startingNode) {
        let visited = new Set();
        let queue = [];

        visited.add(startingNode);
        queue.push(startingNode);

        while (queue.length > 0) {
            let node = queue.shift();
            console.log(node);

            let neighbors = this.adjList.get(node);
            for (let neighbor of neighbors) {
                if (!visited.has(neighbor)) {
                    visited.add(neighbor);
                    queue.push(neighbor);
                }
            }
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
graph.bfs("A");
// 出力:
// A
// B
// C
```

## 最短経路アルゴリズム

グラフ内の 2 点間の最短経路を求めるアルゴリズムです。代表的なアルゴリズムにダイクストラ法があります。

### サンプルコード：ダイクストラ法

```javascript
class Graph {
    constructor() {
        this.adjList = new Map();
    }

    addVertex(v) {
        this.adjList.set(v, []);
    }

    addEdge(v, w, weight) {
        this.adjList.get(v).push({ node: w, weight: weight });
        this.adjList.get(w).push({ node: v, weight: weight }); // 無向グラフの場合
    }

    dijkstra(startNode) {
        let distances = {};
        let pq = new PriorityQueue();
        let previous = {};

        for (let key of this.adjList.keys()) {
            if (key === startNode) {
                distances[key] = 0;
                pq.enqueue(key, 0);
            } else {
                distances[key] = Infinity;
                pq.enqueue(key, Infinity);
            }
            previous[key] = null;
        }

        while (!pq.isEmpty()) {
            let minNode = pq.dequeue().element;
            let neighbors = this.adjList.get(minNode);

            for (let neighbor of neighbors) {
                let alt = distances[minNode] + neighbor.weight;
                if (alt < distances[neighbor.node]) {
                    distances[neighbor.node] = alt;
                    previous[neighbor.node] = minNode;
                    pq.enqueue(neighbor.node, alt);
                }
            }
        }

        return distances;
    }
}

class PriorityQueue {
    constructor() {
        this.items = [];
    }

    enqueue(element, priority) {
        let queueElement = { element, priority };
        let added = false;
        for (let i = 0; i < this.items.length; i++) {
            if (this.items[i].priority > queueElement.priority) {
                this.items.splice(i, 0, queueElement);
                added = true;
                break;
            }
        }
        if (!added) {
            this.items.push(queueElement);
        }
    }

    dequeue() {
        if (this.isEmpty()) return "Underflow";
        return this.items.shift();
    }

    isEmpty() {
        return this.items.length === 0;
    }
}

let graph = new Graph();
graph.addVertex("A");
graph.addVertex("B");
graph.addVertex("C");
graph.addVertex("D");
graph.addEdge("A", "B", 1);
graph.addEdge("A", "C", 4);
graph.addEdge("B", "C", 2);
graph.addEdge("B", "D", 5);
graph.addEdge("C", "D", 1);
let distances = graph.dijkstra("A");
console.log(distances);
// 出力:
// { A: 0, B: 1, C: 3, D: 4 }
```

## 結論

グラフアルゴリズムは、複雑なネットワーク構造の問題を解決するために不可欠です。深さ優先探索（DFS）、幅優先探索（BFS）、ダイクストラ法などの基本アルゴリズムを理解し、適用することで、効率的な問題解決が可能になります。
