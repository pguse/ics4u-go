### A Pluto.jl notebook ###
# v0.12.5

using Markdown
using InteractiveUtils

# ╔═╡ 7aa44100-1d31-11eb-0959-8581ce4e3568
md"# Structs and Objects

At this point we have used functions to organize code and built-in types to organize data.  Now, we are going to build our own types to organize both code and data.

## Composite Types

We are going to build a **composite type** that describes a point in 2-dimensional space. In Julia, a programmer described composite type is called a **struct**.  A struct is like a _blueprint_ that is used for creating **objects** of a particular type.  A **struct** definition for a point would look like this"

# ╔═╡ fb4f18c0-1d31-11eb-3422-63f9e7f67e41
struct Point
	x
	y
end

# ╔═╡ 0c2f7860-1d32-11eb-125d-47b797ba24d6
md"The point has two **attributes** or **fields**:  x and y.  To create a point **object**, we would use what is called the **constructor** function like this"

# ╔═╡ bae4f6f2-1d32-11eb-24fc-0769904fc94e
p1 = Point(3,4)

# ╔═╡ ea172dd0-1d32-11eb-2a82-39a0ef837252
md"Creating a new object is called **instantiation**, and the object is an **instance** of the type.  To access the fields of a composite type, we use the **dot operator**."

# ╔═╡ dbd11630-1d34-11eb-2870-f97f4394a3f8
p1.x

# ╔═╡ fda86ab0-1d34-11eb-3e19-ab135d69c396
md"To use a composite type within a function, it must be passed in as a **parameter**, as any other type.  For example"

# ╔═╡ 59391f02-1d35-11eb-004b-ff17108a1c40
function midpoint(p1::Point, p2::Point)
	(p1.x + p2.x)/2 , (p1.y + p2.y)/2
end

# ╔═╡ 91757bc0-1d35-11eb-0283-dbb5e6d28ed5
md"Notice in the definition of the **midpoint()** function, the **type** of each parameter has been specified.  This is optional, but is a clear way of indicating that this function is designed specifically to be used with our new **Point** type.  Here is an example of how the **midpoint()** function can be used."

# ╔═╡ f9fc3620-1d35-11eb-001c-97c4be3bc7d7
begin
	p2 = Point(8,16)
	midpoint(p1,p2)
end

# ╔═╡ 95bfb320-1d36-11eb-05d9-b1f46ebf52c8
md"## Exercises

1.  Create a function called **distance()** that takes two Point objects as parameters and returns the distance between them."

# ╔═╡ dd22ffb0-1d36-11eb-36c0-53ae1a6c0d9f
function distance(p1::Point, p2::Point)
	(p2.y + p1.y) + (p2.x +p1.x)
end

# ╔═╡ eda0a5e0-1d36-11eb-03d8-a545bbd9fc06
distance( Point(1,1), Point(5,5) )

# ╔═╡ ddb97350-1d36-11eb-14f5-99474aa7815b
md"2.  Create a function called **slope()** that takes two Point objects as parameters and returns the slope of the line segment between the two points."

# ╔═╡ e21ccd70-1d36-11eb-3acb-2d67d23dde06
function slope(p1::Point, p2::Point)
	(p2.y - p1.y) / (p2.x - p1.x)
end

# ╔═╡ 9b5c8ff0-1d3c-11eb-05e3-b573685e09d2
slope( Point(1,1), Point(5,5) )

# ╔═╡ 451b71c0-1d45-11eb-1c03-43838502e192
slope(p1,p2)

# ╔═╡ 08868db0-1cbb-11eb-09da-93aca6bc5f88
md"# Modelling Cards"

# ╔═╡ d8394210-1d37-11eb-2080-3929704be080
md"We are now going to create a type that will be used to model a playing card.  As you can see below, both the **rank** and **suit** fields have been chosen to be integers.  This clearly communicates what types are expected when a **Card** object is created with the **constructor** function."

# ╔═╡ 14013e60-1cbb-11eb-248f-1de694433cad
struct Card
	rank :: Int64
	suit :: Int64
end

# ╔═╡ 2f3e5b90-1cbb-11eb-3f14-adeeb76d1361
c = Card(1,4)

# ╔═╡ 20891180-1d38-11eb-22b5-b3a5371e8914
md"The choice of **integer** fields was made because it is anticipated that we will need to decide the value of each card in the future, and it may be easier to do this with integer values.  However, this is a **choice** in how the **Card** object is designed and other choices could certainly be made.   Again, the **dot operator** is used to access the fields."

# ╔═╡ 3a6dd270-1cbb-11eb-294e-49afabad8469
c.rank

# ╔═╡ 9bb66c90-1d38-11eb-1b48-d7a62b78fd27
md"In order to print Card objects in a way that people can easily read _(since the **fields** are integers)_, we need a mapping from the integer codes to the corresponding ranks and suits. A natural way to do that is with arrays of strings:"

# ╔═╡ 859c7ee0-1cbb-11eb-1f70-c17e4342b7d6
begin
	const rank_names = ["A","2","3","4","5","6","7","8","9","10","J","Q","K"]
	const suit_names = ['♣','♢','♡','♠'] 
end

# ╔═╡ f82f36f0-1d38-11eb-30cb-159f98fa0dc2
md"The variables **suit\_names** and **rank\_names** are **global variables**. The **const** declaration means that the variable can only be assigned once.  For every built-in type, Julia provides a **show()** function that determine how the type should be displayed.  Since we have created a new **type** called **Card**, we can create our own **show()** function that displays a **Card**."

# ╔═╡ 73c88520-1d37-11eb-0f6b-c188e53fcb5c
function Base.show(io::IO, card::Card)
    print(io, rank_names[card.rank], suit_names[card.suit])
end

# ╔═╡ f48bb4f0-1d39-11eb-3922-35c16aa23a10
md"Are card is displayed as follows.  Notice, how the **const** arrays created above have been used."

# ╔═╡ 86bc8282-1d37-11eb-0d92-4b6159cd4263
c

# ╔═╡ 4130d4c0-1d3a-11eb-2100-cb2322ffac30
md"Here is a snippet of code that generates a **random** array of **Cards**."

# ╔═╡ 73ef8dc0-1d3a-11eb-07f5-4f5a0650ebf3
hand = [ Card(1,1), Card(2,1), Card(3,1), Card(4,1), Card(5,1) ]

# ╔═╡ 9d4b4092-1d3c-11eb-37ca-392db675558e
md"## Exercises:

1. Create a **Boolean** function called **flush(h)** that takes in an array of cards and determines whether it is a flush _(all the same suit)_."

# ╔═╡ db923ed0-1d3c-11eb-10c0-4f0f921dafbe
function flush(h::Array{Card,1})
	h[1].suit == h[2].suit == h[3].suit == h[4].suit == h[5].suit
end

# ╔═╡ e0e269f2-1d3c-11eb-227a-6bfe9d399e6a
flush(hand::Array{Card,1})

# ╔═╡ c0e46950-1d3c-11eb-3151-15159ba46107
md"2. Create a **Boolean** function called **straight(h)** that takes in a hand of cards and determines whether it is a straight _(sequence of ranks)_.  _**Note:**  Assume the hand of cards is already sorted_."

# ╔═╡ f3683c30-1d3c-11eb-399f-09bcd47c8508
function straight(h::Array{Card,1})
	h[1].rank == h[2].rank - 1 && h[2].rank == h[3].rank - 1 && h[3].rank == h[4].rank - 1 && h[4].rank == h[5].rank - 1
end

# ╔═╡ 3964a5c0-1d3d-11eb-1e4b-e7b70fce91e3
straight(hand::Array{Card,1})

# ╔═╡ 0900f3be-1d3d-11eb-0f4d-9baa90770627
md"3. Create a **Boolean** function called **straightFlush(h)** that takes in a hand of cards and determines whether it is a straight flush _(sequence of ranks, all the same suit)_.  _**Note:**  Assume the hand of cards is already sorted_."

# ╔═╡ fb5cac50-1d3c-11eb-1517-7ba52ae11740
function straightFlush(h::Array{Card,1})
	h[1].rank == h[2].rank - 1 && h[2].rank == h[3].rank - 1 && h[3].rank == h[4].rank - 1 && h[4].rank == h[5].rank - 1 && h[1].suit == h[2].suit == h[3].suit == h[4].suit == h[5].suit
end

# ╔═╡ 38760000-1d3d-11eb-1fe1-67404044cfa9
straightFlush(hand::Array{Card,1})

# ╔═╡ Cell order:
# ╟─7aa44100-1d31-11eb-0959-8581ce4e3568
# ╠═fb4f18c0-1d31-11eb-3422-63f9e7f67e41
# ╟─0c2f7860-1d32-11eb-125d-47b797ba24d6
# ╠═bae4f6f2-1d32-11eb-24fc-0769904fc94e
# ╟─ea172dd0-1d32-11eb-2a82-39a0ef837252
# ╠═dbd11630-1d34-11eb-2870-f97f4394a3f8
# ╟─fda86ab0-1d34-11eb-3e19-ab135d69c396
# ╠═59391f02-1d35-11eb-004b-ff17108a1c40
# ╟─91757bc0-1d35-11eb-0283-dbb5e6d28ed5
# ╠═f9fc3620-1d35-11eb-001c-97c4be3bc7d7
# ╟─95bfb320-1d36-11eb-05d9-b1f46ebf52c8
# ╠═dd22ffb0-1d36-11eb-36c0-53ae1a6c0d9f
# ╠═eda0a5e0-1d36-11eb-03d8-a545bbd9fc06
# ╟─ddb97350-1d36-11eb-14f5-99474aa7815b
# ╠═e21ccd70-1d36-11eb-3acb-2d67d23dde06
# ╠═9b5c8ff0-1d3c-11eb-05e3-b573685e09d2
# ╠═451b71c0-1d45-11eb-1c03-43838502e192
# ╟─08868db0-1cbb-11eb-09da-93aca6bc5f88
# ╟─d8394210-1d37-11eb-2080-3929704be080
# ╠═14013e60-1cbb-11eb-248f-1de694433cad
# ╠═2f3e5b90-1cbb-11eb-3f14-adeeb76d1361
# ╟─20891180-1d38-11eb-22b5-b3a5371e8914
# ╠═3a6dd270-1cbb-11eb-294e-49afabad8469
# ╟─9bb66c90-1d38-11eb-1b48-d7a62b78fd27
# ╠═859c7ee0-1cbb-11eb-1f70-c17e4342b7d6
# ╟─f82f36f0-1d38-11eb-30cb-159f98fa0dc2
# ╠═73c88520-1d37-11eb-0f6b-c188e53fcb5c
# ╟─f48bb4f0-1d39-11eb-3922-35c16aa23a10
# ╠═86bc8282-1d37-11eb-0d92-4b6159cd4263
# ╟─4130d4c0-1d3a-11eb-2100-cb2322ffac30
# ╠═73ef8dc0-1d3a-11eb-07f5-4f5a0650ebf3
# ╟─9d4b4092-1d3c-11eb-37ca-392db675558e
# ╠═db923ed0-1d3c-11eb-10c0-4f0f921dafbe
# ╠═e0e269f2-1d3c-11eb-227a-6bfe9d399e6a
# ╟─c0e46950-1d3c-11eb-3151-15159ba46107
# ╠═f3683c30-1d3c-11eb-399f-09bcd47c8508
# ╠═3964a5c0-1d3d-11eb-1e4b-e7b70fce91e3
# ╟─0900f3be-1d3d-11eb-0f4d-9baa90770627
# ╠═fb5cac50-1d3c-11eb-1517-7ba52ae11740
# ╠═38760000-1d3d-11eb-1fe1-67404044cfa9
