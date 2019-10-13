package app


//func RandToken(l int) string {
//	b := make([]byte, l)
//	rand.Read(b)
//	return base64.StdEncoding.EncodeToString(b)
//}
//
//func GetToken()  {
//
//	store := sessions.NewCookieStore([]byte(handlers.RandToken(64)))
//	store.Options(sessions.Options{
//		Path:   "/",
//		MaxAge: 86400 * 7,
//	})
//}