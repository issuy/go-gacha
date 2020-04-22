# はじめに
（社内勉強会でやったネタ）

## きっかけと、なにをやるか
- Goを久しぶりに触ってみよう
    - Go完全ににわか（ノリとテンションでやってる）
    - ちょうど社内のGo勉強会で基礎を学べる機会があったしある程度モチベある
- ガチャっぽい実装にするか
    - 軽く触るのでお題は適当でいいや
    - とはいえ経験のないお題だと考えること増えるので今回やりたいこととはズレる
    - 趣旨的には経験のあるお題が良さそう
    - 復習がてらガチャっぽいのやるか

→　**よっしゃ久しぶりにガチャっぽいのでも実装してみよう**

## このテキストで触れる内容
1. 作業手順
    - どのような手順で実装していったか、コミット単位で追っていく
1. 解説となぜその実装にしたか？
    - 自分の復習のために「こんなのあったな」を思い出して記載
    - 「理由が大事」という意識があるので思想や裏側の共有

## レビュー歓迎
https://github.com/issuy/go-gacha/pull/1
気になること・間違いなどあればガンガン指摘お願いします！
勘違い・見落としなど絶対ある。。。
[追記]実際書きながらコード読み返してたらちらほら。。。

## 実行環境
Go 1.14.2

# よっしゃやってこ
コミット単位で追っていく

## [Define procedure.](https://github.com/issuy/go-gacha/pull/1/commits/cac20075e6a3fa9e478ab6d67d06edbc714389a3)
### まずは手続きを決める。
- [Information]
  - 確率表記とか排出アイテムの出力
- [Execute]
  - 抽選処理を実行
- [Result]
  - 抽選結果の出力

## [Definition of rarity and probability.](https://github.com/issuy/go-gacha/pull/1/commits/8794011ca2fc012ebbd8e6a60b73ff87e65ec852)
### Rarityの定義
idはDBっぽくしてるだけ。使わないかも。
rateはそのレアリティの排出確率。万分率で表現。
### 万分率
万分率における `1` は割合における `0.0001` で `1/10000` の確率を表現する。
つまり `100` は `1%` になる。
※なぜ万分率使ってるかは後で説明

## [Create a probability calculator.](https://github.com/issuy/go-gacha/pull/1/commits/581ee2b46c99166c37c3ef175ca2080c85d03001)
### ProbabilityCalculatorの定義
確率計算機。
denominator には万分率で定義した確率の sum が入る。

### GetRate()の定義
各レアリティの抽選確率を文字列で返してくれる。

Goでいうレシーバを使ってるがふんわりした理解。
またGo勉強会で触れてくれると思うので一旦置いておく。

レシーバを使うことで、
```
calculator := GetProbabilityCalculator(rarities)
calculator.GetRate(rarity)
```
みたいなメソッドの使い方ができるようになる。

### 万分率使う理由 その１
GetRate()では、たとえば rete=6000 の denominator=10000 で `60.0000%` という文字列形式になってる。
運用的にきれいな数字をいれておかないと、 rete=10 の denominator=30 で `33.33333333...%` とかなっちゃう。
この場合小数点以下n位切り捨てとか、四捨五入にするかとか考えなければならなくなる。
**ちゃんとしないと「見え方」としての確率と「抽選ロジック」としての確率がズレる。**

万分率で万の範囲の割合を表現する決まりを作っておけば、そこに困らなくなる。
例えば[こんな感じ](https://play.golang.org/p/Guq7EK2crmQ)で  `3333` は `33.33%` になる。

その代わり小数点以下の表現範囲が制限される。
より広範囲を表現したければ100万分率など広げていけばいい。

## [Output probability for each rarity.](https://github.com/issuy/go-gacha/pull/1/commits/711e81fd6635ee357c6f7fc7806d341434a61009)
### 出力結果
```
[Information]
Rarity & Rate(%):
R=60.0000 SR=35.0000 UR=5.0000 
Rarity & Item:
[Execute]
Draw
[Result]
Rarity:---
Item:---
```

## [Add items and associate with rarity.](https://github.com/issuy/go-gacha/pull/1/commits/33dc8d9dc6221f1ae0eda554253442888d3f3c0f)
### Itemの定義とRarityへの追加
Rarity と Item の紐付けをどう持つかな〜？と思ったけど、ここでは処理しやすいように Rarity に追加。
Rails の Association っぽい構成にした。

### GetRarities()にItemのデータを定義
地味に時間がかかったのがここ。（良い感じのアイテム名が。。。
データの定義は雑に1個のArrayで定義。なんとなく名前でレアリティが分かるようにしておく。（俺にはわかる

それをSliceしてレアリティごとに振り分け。

## [Output items for each rarity.](https://github.com/issuy/go-gacha/pull/1/commits/eed849a3387c16cd9fdada1a02d826814068b8d5)
### 出力結果
```
[Information]
Rarity & Rate(%):
R=60.0000 SR=35.0000 UR=5.0000 
Rarity & Item:
R=[ きれいな石ころ イイカンジの枝 ドライバー ネジ 泥水 真水 付箋 定規 土 草 ]
SR=[ イケてるTシャツ ビール缶6本セット チョコレートアソート スタバカード 医療用マスク詰め合わせ ちょっといいぬいぐるみ 1000円分の商品券 ]
UR=[ 金の延べ棒 ダイヤの指輪 ディズニーペアチケット ]

[Execute]
Draw
[Result]
Rarity:---
Item:---
```

## [Draw rarity.](https://github.com/issuy/go-gacha/pull/1/commits/3a6f6d33ee4698b52798843e115ca11d4160b733)
### 抽選ロジック
まずレアリティ、次にそのレアリティに紐づくアイテム、という順番で抽選していく。
まずはレアリティの抽選から。

### なぜこの順番なのか
**`R=60% SR=35% UR=5%` みたいな表記**を行った場合、抽選結果は**必ずそれに準ずる確率で排出**されないとダメ。
さらにその**レアリティの中から抽選されるアイテム**は、**個別の確率表記がなければ普通は当倍率で抽選**される。

例えばその１：アイテムのデータ定義時にレアリティに応じた確率を設定するケース。
数が多いので大変。設定ミスると障害だし。
```
ItemA=12.34%
ItemB=7.44%
...
```

例えばその２：抽選ロジックで各アイテムごとの確率を計算してから1つを選ぶケース。
R=60%の中に7個のアイテムがあるとしたら一つあたりの確率は `8.571428571428571%` になるが、こうなってくると浮動小数点として扱わなければならなくなる。
浮動小数点の計算をするなら誤差も考慮しないといけなくなるので注意が必要。

誤差の話は後（万分率使う理由 その２）で説明するが、できるだけ整数値で扱って上げるのが安全。

### レアリティ抽選のしかた
乱数生成して、それを dart(矢) とする。
それを board に刺すが、 board はレアリティごとの rate に応じた専有面積をもつ。
dart が刺さった敷地がどのレアリティだったかで抽選結果を決める。

文字で表現するならこう↓
```
確率は各10%で定義。乱数生成機は0~9を生成する場合。

R == a (60%)
SR == b (30%)
UR == c (10%)

[aaaaaa|bbb|c]
[012345|678|9]

抽選の例：
R=6
SR=3
UR=1
dart = 6
board = 0

board += R
dart < board ならRを排出
dart=6, board=6なので false （dart は0始まりなので注意）
board += SR
dart < board ならSRを排出
dart=6, board=9なので true
抽選の結果、SRを排出
```
これでレアリティの抽選は完了！

### 乱数生成機：メルセンヌ・ツイスタ
横着して[記事](https://makiuchi-d.github.io/2017/09/09/qiita-9c4af327bc8502cdcdce.ja.html)から拝借した。
一言でいうと「よりイイカンジの乱数生成器」。

### 閉区間・開区間
[rand. Int63n(n)](https://golang.org/pkg/math/rand/#Rand.Int63n) のドキュメントに 
```
Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
```
とある。
これを見た時に「あ〜久しぶりに見た」と思ったのでついでに残しておく。

この `[0, n)` は数学的な 「区間」を表すもの。
`[` 側を「閉区間」、`)` 側を「開区間」と呼ぶ。
`[0, max)` は乱数生成機から得られる数値をxとすると、 `0 <= x < n` のようなxが得られる。

ちなみに僕は大学のころ数学のテストで赤点とりました。

### [rand.Int63n()](https://golang.org/pkg/math/rand/#Int63n)の引数はInt64
`Int63n()` に入れる予定だった denominator を動的型付けの癖で適当に `int` で宣言していた。
ここ整理しようと思ったけど力尽きたごめんなさい。

## [Output draw rarity result.](https://github.com/issuy/go-gacha/pull/1/commits/dd35074effec6177d591b43a891e9f687a3ad907)
### 出力結果
```
[Information]
Rarity & Rate(%):
R=60.0000 SR=35.0000 UR=5.0000 
Rarity & Item:
R=[ きれいな石ころ イイカンジの枝 ドライバー ネジ 泥水 真水 付箋 定規 土 草 ]
SR=[ イケてるTシャツ ビール缶6本セット チョコレートアソート スタバカード 医療用マスク詰め合わせ ちょっといいぬいぐるみ 1000円分の商品券 ]
UR=[ 金の延べ棒 ダイヤの指輪 ディズニーペアチケット ]

[Execute]
Draw...
Debug:R 6000 6000 6596
Debug:SR 3500 9500 6596

[Result]
Rarity:SR
Item:---
```
レアリティが決まったゾ！（意外と良いの出た

## [Cut out the rand process.](https://github.com/issuy/go-gacha/pull/1/commits/47fd99649c7336cb63c3cb26fc894936dea6e290)
### 乱数生成期の初期化処理をメソッドに切り出した
この次に実装する処理で使おうと思ったので切り出して共通化した。

## [Draw item.](https://github.com/issuy/go-gacha/pull/1/commits/5dd43edd71bfc371c24205851ad7703e7b4d78fd)
### アイテム抽選のしかた
今回はアイテムは等倍率で抽選。
簡単にレアリティとは違う抽選ロジックにしちゃう。

文字で表現するならこう↓
```
要素数が３の配列の中で、どのインデックスの要素を排出するか？
この場合は乱数生成機から 0~2 の数値を取得すれば良い。
得られた数値が 0 なら array[0] を、1 なら array[1] を排出する。
```
これでアイテムの抽選は完了！

### 万分率使う理由 その２
前述のようにレアリティ・アイテムの**抽選はすべて整数値で処理**している。
これができるように万分率の決まりに従ってデータ定義している。

なぜこんなことをするかというと、**浮動小数点を扱うと誤差がでるから**。
[こんな簡単な足し算](https://play.golang.org/p/glzFiowHE07)だけでも誤差が生じる。
それを避け安全に計算するために整数値で扱う。

## [Output draw item result.](https://github.com/issuy/go-gacha/pull/1/commits/f9186599de53914a40bfebe0916aa06b637363fe)
### 出力結果
```
[Information]
Rarity & Rate(%):
R=60.0000 SR=35.0000 UR=5.0000 
Rarity & Item:
R=[ きれいな石ころ イイカンジの枝 ドライバー ネジ 泥水 真水 付箋 定規 土 草 ]
SR=[ イケてるTシャツ ビール缶6本セット チョコレートアソート スタバカード 医療用マスク詰め合わせ ちょっといいぬいぐるみ 1000円分の商品券 ]
UR=[ 金の延べ棒 ダイヤの指輪 ディズニーペアチケット ]

[Execute]
Draw...
Debug:R 6000 6000 6841
Debug:SR 3500 9500 6841
Debug:5

[Result]
Rarity:SR
Item:ちょっといいぬいぐるみ
```
ちゃんとレアリティに応じたアイテムが排出された！（運も良い

## [Refactor the order of methods.](https://github.com/issuy/go-gacha/pull/1/commits/8559e1359f335d23626b085042c0a5228a2175e9)
### 並び順が気になったので位置変えた
が、これで良かったかな？前よりは良くなった感。




