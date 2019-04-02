- Vertex
    - (v *Vertex) method() v.method, (&v).method
    - method(v *Vertex) method(&v)
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
