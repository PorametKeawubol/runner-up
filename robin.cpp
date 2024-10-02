#include <iostream>
#include <vector>
using namespace std;

int main()
{
    int t;
    cin >> t;

    while (t > 0)
    {
        int kon, kuy, sgold, ckon;
        cin >> kon >> kuy;
        vector<int> g(kon);
        for (int i = 0; i < kon; i++)
        {
            cin >> g[i];
        }
        sgold = 0;
        ckon = 0;
        for (auto gkon : g)
        {
            if (gkon >= kuy)
            {
                sgold += gkon;
            }
            else if (sgold >= 1 && gkon == 0)
            {
                sgold--;
                ckon++;
            }
        }
        t--;
        cout << ckon << endl;
    }
}