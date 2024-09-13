#include "iostream"
#include <string>
using namespace std;
int main()
{
    int n;
    cin >> n;
    string w;
    while (n > 0)
    {
        cin >> w;
        if (w.length() > 10)
        {
            cout << w[0] << w.length() - 2 << w[w.length() - 1] << endl;
        }
        else
        {
            cout << w << endl;
        }
        n -= 1;
    }
}