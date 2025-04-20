#include "iostream"
using namespace std;

#include "singlyLinkedList.hpp"

int main(int argc, char *argv[]) {
	cout << "TEST: singlyLinkedList\n";

	//Define the template type for the linked list - below is of integer type
	SinglyLinkedList<int> list;

	//Add a node in the begining of the list 
	list.AddNodeToBegin(10);
	list.AddNodeToBegin(20);
	list.AddNodeToEnd(30);
	list.AddNodeToEnd(40);
	list.AddNodeToBegin(50);
	list.AddNodeToEnd(60);
	list.AddNodeToMiddle(70, 1);
	list.AddNodeToMiddle(80, 3);
	list.PrintMyList();


	//List of type string
	SinglyLinkedList<string> slist;

	slist.AddNodeToMiddle("Zeroth", 1);
	slist.AddNodeToBegin("First");
	slist.AddNodeToBegin("Second");
	slist.AddNodeToEnd("Third");
	slist.AddNodeToMiddle("Fourth", 2);
	slist.PrintMyList();
}
