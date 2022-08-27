# このリポジトリの概要
クリーンアーキテクチャの学習目的での写経リポジトリ

元ネタはこちら
https://qiita.com/arkuchy/items/874656b33d2e5acdf281


# アーキテクチャ
## Entity
ドメインロジック。この処理では本当にデータ型のみ定義されており、ロジックはない

## UseCase
Entityのオブジェクトを操作し、ビジネスロジックを実行
ここではさらにportを定義。portとは下記のInterface Adapterを切り替える為のinterface
InputPortは入力、OutputPortは出力する為のport

## Interface Adapters
UseCaseで定義されたinterfaceの実際の実装がこの部分
このレイヤはDB操作・HTTPの入出力を行う
さらにこの中は以下のように別れている

* controller: 入力に関するadapter。今回であればHTTPのリクエストを担当
* presenter: 出力に関するadapter。今回であればHTTPのレスポンスを担当
* gateway: 永続化に関するadapter。今回であれば、userデータ取得部分を担当

## Frameworks & Drivers
DB Connectionの生成・HTTPのroutingなど技術的な実装




# httpリクエストを受け取った時の処理フロー
※元ネタのpackage同士の関係の方がわかりやすいかもしれない。

1. adapter/controller/user.go#GetUserByID() # ここでリクエストパラメータの受取
1. usecase/interactor/user.go#GetUserByID() # 実行制御。下記のDBアクセスとレスポンス返却を制御している
1. adapter/gateway/user.go#GetUserByID() # DBアクセス
1. adapter/presenter/user.go#Render()	 # レスポンス返却


## 階層を粗結合にしているinterfaceの定義場所
usecase/port/user.goにて上記実行メソッドのinterfaceが定義されており、ここで各階層がinterfaceにより分離されている

