package err

type E struct{

}

func(*E) Get(a string) string {
	return "gun"
}

