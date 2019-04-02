#include <iostream>
#include <thread>
#include <vector>
using namespace std;

class Student {
public:
    Student(string name, int age) {
        _name = name;
        _age = age;
    }

    ~Student() = default;

    int study(int courses) {
        cout << _name << "study " << courses << " courses";
    }
private:
    string _name;
    int _age;
};

void thread_fun(int a)
{
    int i = 0;
    while(true) {
        i+= i^2;
    }
}

int main()
{
    vector<std::thread> v_threads;
    for(int i = 0; i < 4; i++) {
        std::thread t(&thread_fun, 12);
        v_threads.push_back(std::move(t));
    }
    
    v_threads[0].join();
    

    Student xiaoming("小明", 12);
    xiaoming.study(4);
    return 0;
}