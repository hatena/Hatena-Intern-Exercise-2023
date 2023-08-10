# TypeScript 事前課題

この事前課題は、以下のことを目的としています。

- TypeScript の言語機能の習得
- ブラウザの DOM API に触れる

## 課題 1: TypeScript の習得

[TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/intro.html) を一通り読みましょう。

## 課題 2: FizzBuzz

TypeScript を使って、ブラウザ上で FizzBuzz が動作するようにしましょう。[`src/fizzbuzz.ts`](./src/fizzbuzz.ts) を編集して、以下の要件が満たされるようにしてください。

FizzBuzz ボタンを押したら、`ol`要素に次々と`li`要素が追加され、FizzBuzz が表示されるようにしてください。

N 番目の`li`要素は、N が 3 で割り切れるとき“Fizz”、5 で割り切れるとき“Buzz”、3 でも 5 で割り切れるとき“FizzBuzz”、それ以外のときは N の数字をテキストノードとして持ちます。

はじめは以下のように空の`ol`要素があります。

```html
<ol></ol>
```

一度ボタンを押したら以下のように、ひとつ`li`要素が追加されます。

```html
<ol>
  <li>1</li>
</ol>
```

何度もボタンを押すと、以下のように FizzBuzz のルールに従って`li`要素が追加されていきます。

```html
<ol>
  <li>1</li>
  <li>2</li>
  <li>Fizz</li>
  <li>4</li>
  <li>Buzz</li>
  <li>6</li>
  ...
</ol>
```

[DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model) と呼ばれる API を利用するとよいでしょう。

### 取り組み方

リポジトリのルートで、以下のようにローカルサーバーを起動できます。

```console
$ make ts-dev
```

この状態で [http://localhost:9000](http://localhost:8080) にアクセスすると、FizzBuzz のページを見ることができます。

期待通りに動作しているかどうかは、リポジトリのルートで、以下のようにテストを実行することで確認できます。

```console
$ make ts-test
```

### ヒント

FizzBuzz を実装するには何通りもの方法があります。

例えば以下の interface を満たすような`FizzBuzz` class を作る方法があるでしょう。

```typescript
interface FizzBuzz {
  next(): string;
}
```

あるいは以下のように[Generator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Generator) にすることもできるでしょう。

```typescript
function* fizzbuzz(): IterableIterator<string> {}
```

自由な発想で FizzBuzz を実現してください。

## 参考資料

- [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/intro.html)
  - TypeScript について一通り学べる
- [MDN Web Docs](https://developer.mozilla.org/en-US/)
  - Web ページを構成する技術について網羅的にまとまっている

### ツール

- [Yarn](https://yarnpkg.com)
- [webpack](https://webpack.js.org)
- [Babel](https://babeljs.io)
- [jest](https://jestjs.io)

### Developer Tools

ブラウザの Developer Tools でログを確認したり、デバッガーを使ったりできます。

- [Chrome DevTools](https://developer.chrome.com/docs/devtools/)
- [Firefox Developer Tools](https://developer.mozilla.org/en-US/docs/Tools)
- [Safari Developer Tools](https://support.apple.com/ja-jp/guide/safari-developer/welcome/mac)
