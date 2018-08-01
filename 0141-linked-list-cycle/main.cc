#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    bool hasCycle(ListNode *head) {
        if (head == NULL || head->next == NULL) {
            return false;
        }

        ListNode * n1 = head;
        ListNode * n2 = head;
        while(n1 != NULL && n2 != NULL) {
            n1 = n1->next;
            if(n1 != NULL) {
                n1 = n1->next;
            } else {
                continue;
            }

            n2 = n2->next;
            if (n1 == n2 && n1 != NULL) {
                return true;
            }
        }

        return false;
    }
};

int main() {
    ListNode *head = new ListNode(1);
    Solution s;
    cout << s.hasCycle(head) << endl;
    return 1;
}