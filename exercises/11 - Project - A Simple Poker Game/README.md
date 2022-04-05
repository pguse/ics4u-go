# A Simple Poker Game

This is an example of a **Go** program that has been organized using multiple files.

##
Here is your task:

Create a **folder** on your computer called **poker**. Copy all of the go files from here:  **main.go**, **card.go**, **hand.go**, and **deck.go**; from here into your **poker** folder.  In **Visual Studio Code** open your **poker** folder, then create a **New Terminal**.

Your **poker** folder should now look like this:

![poker folder](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderTerminal.png)

In order to run our program, we need to do two new steps.  First, we will define our project as a **module**.  We do this in the **Terminal** using the command,

```
go mod init poker
```

Now your **poker** folder should look like this

![poker folder with go.mod file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderModuleTerminal.png)

The **go.mod** file that you have created looks like this when you open it.

![go mod file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderModuleFile.png)

It just contains the name of the module **poker** and the version of **Go** you were using when creating the file.

The second new step involves how we run our project.  In the **Terminal**, when we are within the **poker** folder, type

```
go run .
```

and you should see the following output.

![go run file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderRun.png)


## Exercises

### 11-0: Flush-Straight-StraightFlush-Pair Functions

Go back to the **Struct** exercises and copy your code from your **Flush**, **Straight**, **StraightFlush**, and **Pair** functions into the **hand.go** file.

### 11-1: Three-of-A-Kind

Create a **ThreeOfAKind** function in the **hand.go** file.

```go
func (h Hand) ThreeOfAKind() bool {
	// Three cards with the same rank - assume a sorted hand
	return false
}
```

### 11-2: Four-of-A-Kind

Create a **FourOfAKind** function in the **hand.go** file.

```go
func (h Hand) FourOfAKind() bool {
	// Four cards with the same rank - assume a sorted hand
	return false
}
```

### 11-3: Full House

Create a **FullHouse** function in the **hand.go** file.

```go
func (h Hand) FullHouse() bool {
	// Two cards with the same rank and Three cards with same rank - assume a sorted hand
	return false
}
```