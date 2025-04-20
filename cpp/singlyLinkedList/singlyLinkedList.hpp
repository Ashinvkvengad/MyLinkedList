#ifndef __SINGLE_LL__
#define __SINGLE_LL__

#include <iostream>
using namespace std;


template <typename dType>
class Node {
public:
	dType data;
	Node *next;

        //Constructor to initialize the Node
        Node(dType val, int ID) : data(val), next(NULL) {
		nodeId = ID + 1;
	}

	int GetNodeId() {
		return nodeId;
	}

private:
	int nodeId;
};

template <typename T>
class SinglyLinkedList {
private:
	Node<T> *head;

	//Helps to keep track of ID of the node
	int NodeID;

public:
	//Constructor
	SinglyLinkedList() : head(NULL), NodeID(-1) {};

	Node<T> *NewNode(T data);
	void AddNodeToBegin(T data);
	void AddNodeToEnd(T data);
	void AddNodeToMiddle(T data, int index);
	void PrintMyList();



//DeleteNodeFromList
//DeleteList
//FindDataFromList
//FindNodeFromList

};

template <typename T>
Node<T> *SinglyLinkedList<T>::NewNode(T data) {
	return new Node<T>(data, NodeID);
}

template <typename T>
void SinglyLinkedList<T>::AddNodeToBegin(T data) {
	Node<T> *node = NewNode(data);
	
	node->next = head;
	head = node;
	NodeID++;
}

template <typename T>
void SinglyLinkedList<T>::AddNodeToEnd(T data) {
	Node<T> *node = NewNode(data);
	
	if(NULL == head) {
		head = node;
		NodeID++;
		return;
	}

	Node<T> *ptr = head;
	while(ptr = ptr->next, ptr->next);

	ptr->next = node;
	NodeID++;
}

template <typename T>
void SinglyLinkedList<T>::AddNodeToMiddle(T data, int index) {
        Node<T> *node = NewNode(data);

        if(NULL == head) {
                head = node;
		NodeID++;
                return;
        }

	int idx = 0;
        Node<T> *ptr = head, *prv;
	while (idx < index && NULL != ptr) {
		prv = ptr;
		ptr = ptr->next;
		idx++;
	}

	if (idx == index) {
		node->next = prv->next;
		prv->next  = node;
		NodeID++;
	}
}

template <typename T>
void SinglyLinkedList<T>::PrintMyList() {
	Node<T> *ptr = head;
	while(NULL != ptr) {
		cout << "data = " << ptr->data << "  Node ID = " << ptr->GetNodeId() << "\n";
		ptr = ptr->next;
	}
}

#endif
