
#include <bits/stdc++.h>
using namespace std;

class BankAccount {
private:
    string accountNumber;
    string accountName;
    double balance;

public:
    BankAccount(string accountNumber, string accountName, double balance) {
        this->accountNumber = accountNumber;
        this->accountName = accountName;
        this->balance = balance;
    }

    // deposit method
    void deposit(double amount) {
        if (amount <= 0) {
            throw runtime_error("amount must be positive");
        }
        balance += amount;
    }

    // withdraw method
    void withdraw(double amount) {
        if (amount <= 0 || amount > balance) {
            throw runtime_error("withdraw can not happen due to wrong amount");
        }
        balance -= amount;
    }

    // transfer method
    void transfer(BankAccount &receiver, double amount) {
        if (amount <= 0 || amount > balance) {
            throw runtime_error("transfer can not happen due to wrong amount");
        }
        balance -= amount;
        receiver.balance += amount;
    }

    // print user data
    void display() {
        cout << "Account Number: " << accountNumber << "\n";
        cout << "Account Name: " << accountName << "\n";
        cout << "Balance: " << balance << "\n";
    }
};

int main() {
    BankAccount soyaib("123", "Soyaib", 200000);
    BankAccount zihad("10", "Zihad", 20);

    soyaib.display();
    zihad.display();
}