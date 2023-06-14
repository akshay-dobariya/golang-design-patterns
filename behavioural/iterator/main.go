package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/collection"
	"github.com/akshay-dobariya/golang-design-patterns/behavioural/iterator/objects"
)

func main() {
	users := []objects.User{
		objects.CreateUser("user1", 23),
		objects.CreateUser("user2", 24),
		objects.CreateUser("user3", 78),
		objects.CreateUser("user4", 34),
		objects.CreateUser("user5", 90),
	}
	userCollection := collection.UserCollection{
		Users: &users,
	}
	iterator := userCollection.CreateInterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("%+v\n", user)
	}
	userNil := iterator.GetNext()
	fmt.Printf("userNil %+v\n", userNil)
	iterator.RestartInteration()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("%+v\n", user)
	}
	fmt.Println("==============")

	mobiles := []objects.Mobile{
		objects.CreateMobile("samsung"),
		objects.CreateMobile("apple"),
		objects.CreateMobile("xiomi"),
		objects.CreateMobile("oppo"),
		objects.CreateMobile("vivo"),
	}
	mobileCollection := collection.MobileCollection{
		Mobiles: &mobiles,
	}
	iterator = mobileCollection.CreateInterator()
	for iterator.HasNext() {
		mobile := iterator.GetNext()
		fmt.Printf("%+v\n", mobile)
	}
	mobileNil := iterator.GetNext()
	fmt.Printf("mobileNil %+v\n", mobileNil)
	iterator.RestartInteration()
	for iterator.HasNext() {
		mobile := iterator.GetNext()
		fmt.Printf("%+v\n", mobile)
	}
}
