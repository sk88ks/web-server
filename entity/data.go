package entity

// サンプル
type Friend struct {
	Fname string
}

type Person struct {
	UserName   string
	Emails     []string
	Friends    []Friend
	HasContent string
}

// ユーザーページ
type UserPage struct {
	Self        User   // 自身のユーザー情報
	UserFriends []User // フレンドのユーザー
	HasNext     bool   // フレンドが4件以上いるか

	HasRecommendItems bool   // おすすめしたアイテムの有無
	RecommendItems    []Item // おすすめアイテム
	HasReviewItems    bool   // レビューしたアイテムの有無
	ReviewItems       []Item // レビューしたアイテム
}

// アイテムページ
type ItemPage struct {
	Item              Item   // 表示中のアイテム情報
	HasRecommendUsers bool   // このアイテムをおすすめしてくれたユーザーの有無
	RecommendUsers    []User // このアイテムをおすすめしてくれたユーザー
	HasReviews        bool   // 最近投稿されたレビューがあるか
	Reviews           []Item // 最近投稿されたレビューのアイテム情報
}

// 投稿情報ページ
type PostPage struct {
	PostId        string
	PostDateTime  string // yyyy年mm月dd日 HH:MM の書式
	PostItemScore int
	PostItemState string // 購入済み とか
	Item          Item   // レビューされたのアイテム情報
	User          User   // レビューしたユーザー情報

	LikeUsers []User // いいねしてくれたユーザー情報
	HasNext   bool   // いいねユーザーが4件以上か

	// 以下アイテムページと同じ
	HasRecommendUsers bool   // このアイテムをおすすめしてくれたユーザーの有無
	RecommendUsers    []User // このアイテムをおすすめしてくれたユーザー
	HasReviews        bool   // 最近投稿されたレビューがあるか
	Reviews           []Item // 最近投稿されたレビューのアイテム情報
}
