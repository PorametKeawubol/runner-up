#include "iostream"
using namespace std;
int main()
{
    int n;
    cin >> n;
    while (n > 0)
    {
        int a, b;
        cin >> a >> b;
        if (a > b)
        {
            cout << b << ' ' << a;
        }
        else
        {
            cout << a << ' ' << b;
        }
        n -= 1;
    }
}