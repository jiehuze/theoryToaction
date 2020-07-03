package observer

func Test()  {
	var reader1 = NewReader("reader1")
	var reader2 = NewReader("reader2")
	var reader3 =  NewReader("reader3")

	var subject = NewSubject()
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.Notifycation()
}
