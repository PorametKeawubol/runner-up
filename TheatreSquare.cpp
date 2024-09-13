#include "iostream"
#include <cmath>
using namespace std;

int main()
{
    double n, m, a;
    cin >> n >> m >> a;
    double area = n * m, stone = a * a;
    long long ans = ceil((double)n / a) * ceil((double)m / a);

    if (a == 1)
    {
        cout << (long long)area;
    }
    else if (area / stone <= 1)
    {
        cout << 1;
    }
    else if (area / stone > 1)
    {
        cout << ans;
    }
}