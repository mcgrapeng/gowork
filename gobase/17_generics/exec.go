package genericsZhang

import "fmt"

func Exec() {

	fmt.Println(`
	1.【泛型】
		(1).定义：
			Go语言中的【函数】和【类型】支持添加类型参数。
			类型参数列表看起来像普通的参数列表，只不过它使用方括号（[]）而不是圆括号（()）。

			举个例子：
				func min[T int | float64](a, b T) T {
					if a <= b {
						return a
					}
					return b
				}

			其中：
				[T int | float64]为类型参数列表

		(2).注意：
			要使用泛型，必须先进行实例化。

		(3).实例化：
			向函数提供类型参数称为实例化（ instantiation ）。
			1️⃣步骤
				类型实例化分两步进行：
					首先，编译器在整个泛型函数或类型中将所有类型形参（type parameters）替换为它们各自的类型实参（type arguments）。
					其次，编译器验证每个类型参数是否满足相应的约束。

				举个例子：
					fmin := min[float64] // 类型实例化，编译器生成T=float64的min函数
			
			2️⃣调用
				在成功实例化之后，我们将得到一个非泛型函数，它可以像任何其他函数一样被调用。

				举个例子：
					m2 = fmin(1.2, 2.3)  // 1.2

			3️⃣一步到位
				
				举个例子：
					m1 := min[int](1, 2)  // 1
					m2 := min[float64](-0.1, -0.2)  // -0.2

		(4).函数中使用泛型
			1️⃣格式：
				func 函数名称[T 类型参数列表](参数列表 T) T {
					
				}

				举个例子：
					func min[T int | float64](a, b T) T {
						if a <= b {
							return a
						}
						return b
					}

			2️⃣实例化：
				举个例子：
					m1 := min[int](1, 2)  // 1
					m2 := min[float64](-0.1, -0.2)  // -0.2

				这次定义的min函数就同时支持int和float64两种类型，也就是说当调用min函数时，我们既可以传入int类型的参数，
				也可以传入float64类型的参数。

		(5).类型中使用泛型
			1️⃣格式：
				type 类型名称[T 类型参数列表] T
			
				举个例子：
					type Slice[T int | string] []T
	
					type Map[K int | string, V float32 | float64] map[K]V
					
					type Tree[T interface{}] struct {
						left, right *Tree[T]
						value       T
					}
		
			2️⃣实例化：
				举个例子：
					var stringTree Tree[string] //Tree[string]是使用类型实参string实例化 Tree

		(6).类型集
			1️⃣定义：
				类似于参数列表中每个参数都有对应的参数类型，类型参数列表中每个类型参数都有一个类型约束。
				类型约束定义了一个类型集——只有在这个类型集中的类型才能用作类型实参。

			2️⃣格式：
				type V interface {
					类型列表
				}

				举个例子：
					// 类型约束字面量，通常外层interface{}可省略
					func min[T interface{ int | float64 }](a, b T) T {
						if a <= b {
							return a
						}
						return b
					}

					或

					// 作为类型约束使用的接口类型可以事先定义并支持复用
					type Value interface {
						int | float64
					}
					func min[T Value](a, b T) T {
						if a <= b {
							return a
						}
						return b
					}

					或

					type IntPtrSlice[T interface{ *int }] []T

			3️⃣符号：
				1.种类：
					｜ 和 ～
				
				2.解释：
					T1 | T2表示类型约束为T1和T2这两个类型的并集。
						举个例子：
							type Integer interface {
								Signed | Unsigned
							}
						Integer类型表示由Signed和Unsigned组成。

					~T表示所以底层类型是T的类型。（~符号后面只能是基本类型）
						举个例子：
							type Signed interface {
								~int | ~int8 | ~int16 | ~int32 | ~int64
							}
		
		(7).类型推断
			1️⃣函数参数类型推断：
				func min[T int | float64](a, b T) T {
					if a <= b {
						return a
					}
					return b
				}

				显示调用：
					var a, b, m float64
					m = min[float64](a, b) // 显式指定类型实参

				类型推断：（在许多情况下，编译器可以从普通参数推断 T 的类型实参。这使得代码更短，同时保持清晰。）
					var a, b, m float64
					m = min(a, b) // 无需指定类型实参

				注意⚠️：
					这种从实参的类型推断出函数的类型实参的推断称为函数实参类型推断。
					函数实参类型推断只适用于函数参数中使用的类型参数，而不适用于仅在函数结果中或仅在函数体中使用的类型参数。
					例如，它不适用于像 MakeT [ T any ]() T 这样的函数，因为它只使用 T 表示结果。

			2️⃣约束类型推断：
				func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
					r := make(S, len(s))
					for i, v := range s {
						r[i] = v * c
					}
					return r
				}

				编译器推断 E 的类型参数是切片的元素类型的过程称为约束类型推断。
	`)

	fmt.Println("============================================================以下是例子=========================================================")

	fmt.Println(`
		Go 1.18版本增加了对泛型的支持，泛型也是自 Go 语言开源以来所做的最大改变。
	`)

	/**
	泛型语法
	*/
	fmt.Println(`
		泛型为Go语言添加了三个新的重要特性:
			1.函数和类型的类型参数。
			2.将接口类型定义为类型集，包括没有方法的类型。
			3.类型推断，它允许在调用函数时在许多情况下省略类型参数。
	`)
	m1 := min2[int](10, 20)
	m2 := min2[float64](1.1, 1.2)
	fmt.Println(m1, m2)

	fmin1 := min2[float64]
	m3 := fmin1(1.1, 1.2)
	fmt.Println(m3)

	m4 := Slice[int]{}
	m4 = append(m4, 1)
	fmt.Println(m4)
}

func min2[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

func min3[T int](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

//public static <T extend Integer> String min3(T a){
//	if a <= b {
//		return a
//	}
//	return b
//}

func min4[T interface{ int | float64 }](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

type Slice[T int | string] []T
