#include <iostream>
#include <unordered_map>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if(head == NULL) {
            return NULL;
        }
        
        unordered_map<long, int> ptrMap;
        while(head != NULL) {
            if(ptrMap[long(head)]== 1) {
                return head;
            }
            
            // why??????
            // ptrMap.insert(make_pair(long(head), 1));
            ptrMap[long(head)] = 1;
            head = head->next;
        }
        
        return NULL;
    }
};

int main() {
    Solution s;
    ListNode *head = new ListNode(1);
    head->next = head;
    cout << s.detectCycle(head) << endl;
    /*
    unordered_map<long, int> ptrMap;
    ptrMap.insert(make_pair(long(head), 1));
    cout << long(head) <<" " << ptrMap[long(head)] << endl;
    */
}