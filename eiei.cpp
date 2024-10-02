#include <iostream>
using namespace std;

bool eiei(long long n, long long kuy)
{
    long long last = (1LL * n * (n + 1)) / 2;
    long long first = (1LL * (n - kuy) * ((n - kuy) + 1)) / 2;
    long long c = last - first;
    return c % 2 == 0;
}

int main()
{
    int t;
    cin >> t;

    while (t > 0)
    {
        long long n, kuy;
        cin >> n >> kuy;

        if (eiei(n, kuy))
        {
            cout << "YES" << endl;
        }
        else
        {
            cout << "NO" << endl;
        }
        t--;
    }

    return 0;
}
