#include "iostream"
using namespace std;
// int main()
// {
//     int n;
//     cin >> n;
//     for (int c = 0; c < n; c++)
//     {
//         cout << "Hello, World!" << endl;
//     }
// }

// int main()
// {
//     int a, b;
//     cin >> a >> b;
//     do
//     {
//         cout << "$a" << endl;
//         a++;
//     } while (a == b);
// }

string eiei(string n)
{
    return n + "!";
}

int main()
{
    string n;
    cin >> n;
    cout << eiei(n);
}