#include<iostream>
using namespace std;


auto adder() 
{
    auto sum = 0;
    return [=] (int value) mutable {
        sum += value;
        return sum;
    };
}

int main()
{
    auto a = adder();
    int i = 0;
    do
    {
       cout<<"0+...+" <<i <<"= "<< a(i) << endl;
       i++;
    } while (i < 1000000);
    
    return 0;
}