#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

// 定义一个职工类
class Employee {
private:
    int id; // 职工编号
    string name; // 姓名
    int deptId; // 部门编号
    string duty; // 职责
public:
    // 构造函数
    Employee(int id, string name, int deptId, string duty) {
        this->id = id;
        this->name = name;
        this->deptId = deptId;
        this->duty = duty;
    }

    // 获取职工编号
    int getId() {
        return id;
    }

    // 获取姓名
    string getName() {
        return name;
    }

    // 获取部门编号
    int getDeptId() {
        return deptId;
    }

    // 获取职责
    string getduty() {
        return duty;
    }

    // 设置职工编号
    void setId(int id) {
        this->id = id;
    }

    // 设置姓名
    void setName(string name) {
        this->name = name;
    }

    // 设置部门编号
    void setDeptId(int deptId) {
        this->deptId = deptId;
    }

    // 设置职责
    void setduty(string duty) {
        this->duty = duty;
    }

};

// 定义一个管理系统类
class ManagementSystem {
private:
    vector<Employee> employees; // 存储职工信息的向量
    string fileName; // 存储职工信息的文件名

public:
    // 构造函数，从文件中读取职工信息并存入向量中
    ManagementSystem(string fileName) {
        this->fileName = fileName;
        ifstream ifs(fileName); // 打开文件输入流
        if (ifs.is_open()) { // 判断文件是否打开成功
            int id, deptId; 
            string name; 
            string duty; 
            while (ifs >> id >> name >> deptId >> duty) { // 循环读取文件中的每一行数据
                Employee e(id, name, deptId, duty); // 创建一个职工对象
                employees.push_back(e); // 将职工对象添加到向量中
            }
            ifs.close(); // 关闭文件输入流
        }
        else { 
            cout << "文件打开失败！" << endl;
        }

    }

    
    void showMenu() {
        cout << "***************************" << endl;
        cout << "*****欢迎使用管理系统*****" << endl;
        cout << "***** 0.退出管理程序 *****" << endl;
        cout << "***** 1.增加职工信息 *****" << endl;
        cout << "***** 2.显示职工信息 *****" << endl;
        cout << "***** 3.删除离职职工 *****" << endl;
        cout << "***** 4.修改职工信息 *****" << endl;
        cout << "***** 5.按照编号排序 *****" << endl;
        cout << "***** 6.清空所有文档 *****" << endl;
        cout << "***************************" << endl;
        cout << "请输入你的选择：" << endl;
    }

    
    void exitSystem() {
        cout << "感谢使用管理系统，再见！" << endl;
        exit(0);
    }

    // 增加职工信息函数，实现批量添加职工功能,将信息录入到文件中
    void addEmployee() {
        cout << "请输入要添加的职工数量：" << endl;
        int num; // 存储要添加的职工数量
        cin >> num; // 输入职工数量
        if (num > 0) { // 判断输入是否合法
            for (int i = 0; i < num; i++) { // 循环输入每个职工的信息
                int id, deptId; // 存储职工编号和部门编号
                string name; // 存储姓名
                string duty; // 存储职责
                cout << "请输入第" << i + 1 << "个职工的编号：" << endl;
                cin >> id; // 输入职工编号
                cout << "请输入第" << i + 1 << "个职工的姓名：" << endl;
                cin >> name; // 输入姓名
                cout << "请输入第" << i + 1 << "个职工的部门编号：" << endl;
                cin >> deptId; // 输入部门编号
                cout << "请输入第" << i + 1 << "个职工的职责：" << endl;
                cin >> duty; // 输入职责
                Employee e(id, name, deptId, duty); // 创建一个职工对象
                employees.push_back(e); // 将职工对象添加到向量中
            }
            saveFile(); // 将向量中的数据保存到文件中
            cout << "成功添加" << num << "个职工！" << endl;
        }
        else { // 输入不合法，提示错误信息
            cout << "输入有误，请重新输入！" << endl;
        }

    }

    // 显示职工信息函数，显示所有职工的信息
    void showEmployee() {
        if (employees.size() == 0) { // 判断向量是否为空，如果为空，提示没有记录
            cout << "当前没有记录，请先添加！" << endl;
        }
        else { // 如果不为空，打印出所有职工的信息
            cout << "职工编号\t姓名\t部门编号\t职责" << endl;
            for (int i = 0; i < employees.size(); i++) { // 循环遍历向量中的每个元素
                Employee e = employees[i]; // 获取当前职工对象
                cout << e.getId() << "\t\t"
                    << e.getName() << "\t\t"
                    << e.getDeptId() << "\t\t"
                    << e.getduty() << endl; // 打印出当前职工的信息
            }
        }

    }

    // 删除离职职工函数，按照编号删除指定的职工 
    void deleteEmployee() {
        if (employees.size() == 0) { // 判断向量是否为空，如果为空，提示没有记录
            cout << "当前没有记录，请先添加！" << endl;
        }
        else { // 如果不为空，输入要删除的职工编号，并在向量中查找并删除
            cout << "请输入要删除的职工编号：" << endl;
            int id; //职工编号
            cin >> id;
            int index = -1; // 存储要删除的职工在向量中的索引，初始值为-1表示未找到
            for (int i = 0; i < employees.size(); i++) { // 循环遍历向量中的每个元素
                Employee e = employees[i]; // 获取当前职工对象
                if (e.getId() == id) { // 判断当前职工对象的编号是否与要删除的编号相同
                    index = i; // 如果相同，记录当前索引，并跳出循环
                    break;
                }
            }
            if (index != -1) { // 如果找到了要删除的索引，执行删除操作，并保存到文件中，并提示删除成功信息
                employees.erase(employees.begin() + index); // 删除向量中的指定元素
                saveFile(); // 将向量中的数据保存到文件中
                cout << "成功删除职工编号为" << id << "的职工！" << endl;
            }
            else { // 如果没有找到要删除的索引，提示没有找到该职工
                cout << "没有找到职工编号为" << id << "的职工！" << endl;
            }
        }

    }
    // 修改职工信息函数，按照编号修改职工个人信息
    void modifyEmployee() {
        if (employees.size() == 0) { // 判断向量是否为空，如果为空，提示没有记录
            cout << "当前没有记录，请先添加！" << endl;
        }
        else { // 如果不为空，输入要修改的职工编号，并在向量中查找并修改
            cout << "请输入要修改的职工编号：" << endl;
            int id; //职工编号
            cin >> id; 
            int index = -1; // 存储要修改的职工在向量中的索引，初始值为-1表示未找到
            for (int i = 0; i < employees.size(); i++) { // 循环遍历向量中的每个元素
                Employee e = employees[i]; // 获取当前职工对象
                if (e.getId() == id) { // 判断当前职工对象的编号是否与要修改的编号相同
                    index = i; // 如果相同，记录当前索引，并跳出循环
                    break;
                }
            }
            if (index != -1) { // 如果找到了要修改的索引，输入要修改的信息，并更新向量中的元素，并保存到文件中，并提示修改成功信息
                int newId, newDeptId; // 存储新的职工编号和部门编号
                string newName; // 存储新的姓名
                string newduty; // 存储新的职责
                cout << "请输入新的职工编号：" << endl;
                cin >> newId; // 输入新的职工编号
                cout << "请输入新的姓名：" << endl;
                cin >> newName; // 输入新的姓名
                cout << "请输入新的部门编号：" << endl;
                cin >> newDeptId; // 输入新的部门编号
                cout << "请输入新的职责：" << endl;
                cin >> newduty; // 输入新的职责
                Employee e(newId, newName, newDeptId, newduty); // 创建一个新的职工对象
                employees[index] = e; // 更新向量中的指定元素
                saveFile(); // 将向量中的数据保存到文件中
                cout << "成功修改职工编号为" << id << "的职工！" << endl;
            }
            else { // 如果没有找到要修改的索引，提示没有找到该职工
                cout << "没有找到职工编号为" << id << "的职工！" << endl;
            }
        }

    }

    // 按照编号排序函数，按照职工编号，进行排序，排序规则由用户指定
    void sortEmployee() {
        if (employees.size() == 0) { 
            cout << "当前没有记录，请先添加！" << endl;
        }
        else { 
            cout << "请选择排序方式：1.升序 2.降序" << endl;
            int choice; 
            cin >> choice; 
            if (choice == 1) { // 升序排序
                sort(employees.begin(), employees.end(), compareAsc); 
                cout << "按照职工编号升序排序成功！" << endl;
            }
            else if (choice == 2) { // 降序排序
                sort(employees.begin(), employees.end(), compareDesc); 
                cout << "按照职工编号降序排序成功！" << endl;
            }
            else { 
                cout << "输入有误，请重新输入！" << endl;
                return; // 结束函数
            }
            showEmployee(); // 显示排序后的结果
        }

    }

    // 清空所有文档函数，清空文件中记录的所有职工信息 （清空前需要再次确认，防止误删）
    void clearFile() {
        cout << "确认清空？1.确认 2.取消" << endl;
        int choice; // 存储用户选择的确认或取消
        cin >> choice; // 输入用户选择的确认或取消
        if (choice == 1) { // 如果选择确认，打开文件输出流，并清空文件内容，并关闭文件输出流，并清空向量，并提示清空成功信息
            ofstream ofs(fileName, ios::trunc); // 打开文件输出流，使用截断模式
            ofs.close(); // 关闭文件输出流
            employees.clear(); // 清空向量
            cout << "成功清空所有文档！" << endl;
        }
        else { 
            cout << "已取消清空操作！" << endl;
        }

    }

    // 保存文件函数，将向量中的数据保存到文件中
    void saveFile() {
        ofstream ofs(fileName, ios::out); // 打开文件输出流，使用覆盖模式
        for (int i = 0; i < employees.size(); i++) { // 循环遍历向量中的每个元素
            Employee e = employees[i]; // 获取当前职工对象
            ofs << e.getId() << " "
                << e.getName() << " "
                << e.getDeptId() << " "
                << e.getduty() << endl; // 将当前职工的信息写入到文件中
        }
        ofs.close(); // 关闭文件输出流
    }

    // 比较函数，用于升序排序
    static bool compareAsc(Employee e1, Employee e2) {
        return e1.getId() < e2.getId();
    }

    // 比较函数，用于降序排序
    static bool compareDesc(Employee e1, Employee e2) {
        return e1.getId() > e2.getId(); 
    }
};

int main() {
    ManagementSystem ms("employee.txt"); // 创建一个管理系统对象，并传入文件名作为参数
    int choice; 
    while (true) { 
        ms.showMenu(); 
        cin >> choice; 
        switch (choice) { 
        case 0: // 退出
            ms.exitSystem();
            break;
        case 1: // 增加
            ms.addEmployee();
            break;
        case 2: // 显示
            ms.showEmployee();
            break;
        case 3: // 删除离职职工
            ms.deleteEmployee();
            break;
        case 4: // 修改职工信息
            ms.modifyEmployee();
            break;
        case 5: // 按照编号排序
            ms.sortEmployee();
            break;
        case 6: // 清空所有文档
            ms.clearFile();
            break;
        default: // 其他无效的选项，提示错误信息，并重新输入
            cout << "输入有误，请重新输入" << endl;
        }
    }
        return 0;
    }
