#include <iostream>

using namespace std;

int main() {
    string input;
    while (cin >> input && input != "exit") {
        cout << "You said: " << input << endl;
    }
}
