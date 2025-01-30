#include<bits/stdc++.h>
using namespace std;

const int MAX_CREDIT = 500000;

class CreditCard {
private:
    string name;
    double balance;
    double dailyWithdrawal;
    double totalSpent;

public:
    CreditCard(string name, double balance) {
        this->name = name;
        this->balance = balance;
        this->dailyWithdrawal = 0;
        this->totalSpent = 0;
    }

    void withdraw(double amount) {
        if (amount <= 0 || amount > 20000) {
            throw runtime_error("Withdrawal amount must be less than or equal to 20000");
        }
        if (dailyWithdrawal + amount > 100000) {
            throw runtime_error("Daily withdrawal limit exceeded");
        }
        if (balance - amount < 0) {
            throw runtime_error("Insufficient balance");
        }
        dailyWithdrawal += amount;
        balance -= amount;
    }

    void payBill(double amount) {
        if (totalSpent + amount > MAX_CREDIT) {
            throw runtime_error("Spending limit exceeded");
        }
        totalSpent += amount;
        balance -= amount;
    }

    // print user data
    void display() const {
        cout << "Customer name: " << name << "\n";
        cout << "Available Balance: " << balance << "\n";
        cout << "Total Spent: " << totalSpent << "\n";
        cout << "Daily Withdrawal: " << dailyWithdrawal << "\n";
    }
};

int main() {
    CreditCard Soyaib("Soyaib Rahman", MAX_CREDIT);
    Soyaib.withdraw(15000);
    Soyaib.payBill(200000);
    Soyaib.display();

}
