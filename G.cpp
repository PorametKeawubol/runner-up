#include "iostream"
using namespace std;
int main()
{
    int xa, xb, xc, ya, yb, yc, x1, y1, x2, y2, x3, y3;
    cin >> xa >> ya;
    cin >> xb >> yb;
    cin >> xc >> yc;
    x1 = (2 * xa) - x2;
    cout << x1 << " " << yb << endl;

    x2 = (2 * xb) - x3;
    cout << x2 << " " << yb << endl;

    x3 = (2 * xc) - x1;
    cout << x3 << " " << yb << endl;
}