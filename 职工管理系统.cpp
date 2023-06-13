#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
using namespace std;

// ����һ��ְ����
class Employee {
private:
    int id; // ְ�����
    string name; // ����
    int deptId; // ���ű��
    string duty; // ְ��
public:
    // ���캯��
    Employee(int id, string name, int deptId, string duty) {
        this->id = id;
        this->name = name;
        this->deptId = deptId;
        this->duty = duty;
    }

    // ��ȡְ�����
    int getId() {
        return id;
    }

    // ��ȡ����
    string getName() {
        return name;
    }

    // ��ȡ���ű��
    int getDeptId() {
        return deptId;
    }

    // ��ȡְ��
    string getduty() {
        return duty;
    }

    // ����ְ�����
    void setId(int id) {
        this->id = id;
    }

    // ��������
    void setName(string name) {
        this->name = name;
    }

    // ���ò��ű��
    void setDeptId(int deptId) {
        this->deptId = deptId;
    }

    // ����ְ��
    void setduty(string duty) {
        this->duty = duty;
    }

};

// ����һ������ϵͳ��
class ManagementSystem {
private:
    vector<Employee> employees; // �洢ְ����Ϣ������
    string fileName; // �洢ְ����Ϣ���ļ���

public:
    // ���캯�������ļ��ж�ȡְ����Ϣ������������
    ManagementSystem(string fileName) {
        this->fileName = fileName;
        ifstream ifs(fileName); // ���ļ�������
        if (ifs.is_open()) { // �ж��ļ��Ƿ�򿪳ɹ�
            int id, deptId; 
            string name; 
            string duty; 
            while (ifs >> id >> name >> deptId >> duty) { // ѭ����ȡ�ļ��е�ÿһ������
                Employee e(id, name, deptId, duty); // ����һ��ְ������
                employees.push_back(e); // ��ְ��������ӵ�������
            }
            ifs.close(); // �ر��ļ�������
        }
        else { 
            cout << "�ļ���ʧ�ܣ�" << endl;
        }

    }

    
    void showMenu() {
        cout << "***************************" << endl;
        cout << "*****��ӭʹ�ù���ϵͳ*****" << endl;
        cout << "***** 0.�˳�������� *****" << endl;
        cout << "***** 1.����ְ����Ϣ *****" << endl;
        cout << "***** 2.��ʾְ����Ϣ *****" << endl;
        cout << "***** 3.ɾ����ְְ�� *****" << endl;
        cout << "***** 4.�޸�ְ����Ϣ *****" << endl;
        cout << "***** 5.���ձ������ *****" << endl;
        cout << "***** 6.��������ĵ� *****" << endl;
        cout << "***************************" << endl;
        cout << "���������ѡ��" << endl;
    }

    
    void exitSystem() {
        cout << "��лʹ�ù���ϵͳ���ټ���" << endl;
        exit(0);
    }

    // ����ְ����Ϣ������ʵ���������ְ������,����Ϣ¼�뵽�ļ���
    void addEmployee() {
        cout << "������Ҫ��ӵ�ְ��������" << endl;
        int num; // �洢Ҫ��ӵ�ְ������
        cin >> num; // ����ְ������
        if (num > 0) { // �ж������Ƿ�Ϸ�
            for (int i = 0; i < num; i++) { // ѭ������ÿ��ְ������Ϣ
                int id, deptId; // �洢ְ����źͲ��ű��
                string name; // �洢����
                string duty; // �洢ְ��
                cout << "�������" << i + 1 << "��ְ���ı�ţ�" << endl;
                cin >> id; // ����ְ�����
                cout << "�������" << i + 1 << "��ְ����������" << endl;
                cin >> name; // ��������
                cout << "�������" << i + 1 << "��ְ���Ĳ��ű�ţ�" << endl;
                cin >> deptId; // ���벿�ű��
                cout << "�������" << i + 1 << "��ְ����ְ��" << endl;
                cin >> duty; // ����ְ��
                Employee e(id, name, deptId, duty); // ����һ��ְ������
                employees.push_back(e); // ��ְ��������ӵ�������
            }
            saveFile(); // �������е����ݱ��浽�ļ���
            cout << "�ɹ����" << num << "��ְ����" << endl;
        }
        else { // ���벻�Ϸ�����ʾ������Ϣ
            cout << "�����������������룡" << endl;
        }

    }

    // ��ʾְ����Ϣ��������ʾ����ְ������Ϣ
    void showEmployee() {
        if (employees.size() == 0) { // �ж������Ƿ�Ϊ�գ����Ϊ�գ���ʾû�м�¼
            cout << "��ǰû�м�¼��������ӣ�" << endl;
        }
        else { // �����Ϊ�գ���ӡ������ְ������Ϣ
            cout << "ְ�����\t����\t���ű��\tְ��" << endl;
            for (int i = 0; i < employees.size(); i++) { // ѭ�����������е�ÿ��Ԫ��
                Employee e = employees[i]; // ��ȡ��ǰְ������
                cout << e.getId() << "\t\t"
                    << e.getName() << "\t\t"
                    << e.getDeptId() << "\t\t"
                    << e.getduty() << endl; // ��ӡ����ǰְ������Ϣ
            }
        }

    }

    // ɾ����ְְ�����������ձ��ɾ��ָ����ְ�� 
    void deleteEmployee() {
        if (employees.size() == 0) { // �ж������Ƿ�Ϊ�գ����Ϊ�գ���ʾû�м�¼
            cout << "��ǰû�м�¼��������ӣ�" << endl;
        }
        else { // �����Ϊ�գ�����Ҫɾ����ְ����ţ����������в��Ҳ�ɾ��
            cout << "������Ҫɾ����ְ����ţ�" << endl;
            int id; //ְ�����
            cin >> id;
            int index = -1; // �洢Ҫɾ����ְ���������е���������ʼֵΪ-1��ʾδ�ҵ�
            for (int i = 0; i < employees.size(); i++) { // ѭ�����������е�ÿ��Ԫ��
                Employee e = employees[i]; // ��ȡ��ǰְ������
                if (e.getId() == id) { // �жϵ�ǰְ������ı���Ƿ���Ҫɾ���ı����ͬ
                    index = i; // �����ͬ����¼��ǰ������������ѭ��
                    break;
                }
            }
            if (index != -1) { // ����ҵ���Ҫɾ����������ִ��ɾ�������������浽�ļ��У�����ʾɾ���ɹ���Ϣ
                employees.erase(employees.begin() + index); // ɾ�������е�ָ��Ԫ��
                saveFile(); // �������е����ݱ��浽�ļ���
                cout << "�ɹ�ɾ��ְ�����Ϊ" << id << "��ְ����" << endl;
            }
            else { // ���û���ҵ�Ҫɾ������������ʾû���ҵ���ְ��
                cout << "û���ҵ�ְ�����Ϊ" << id << "��ְ����" << endl;
            }
        }

    }
    // �޸�ְ����Ϣ���������ձ���޸�ְ��������Ϣ
    void modifyEmployee() {
        if (employees.size() == 0) { // �ж������Ƿ�Ϊ�գ����Ϊ�գ���ʾû�м�¼
            cout << "��ǰû�м�¼��������ӣ�" << endl;
        }
        else { // �����Ϊ�գ�����Ҫ�޸ĵ�ְ����ţ����������в��Ҳ��޸�
            cout << "������Ҫ�޸ĵ�ְ����ţ�" << endl;
            int id; //ְ�����
            cin >> id; 
            int index = -1; // �洢Ҫ�޸ĵ�ְ���������е���������ʼֵΪ-1��ʾδ�ҵ�
            for (int i = 0; i < employees.size(); i++) { // ѭ�����������е�ÿ��Ԫ��
                Employee e = employees[i]; // ��ȡ��ǰְ������
                if (e.getId() == id) { // �жϵ�ǰְ������ı���Ƿ���Ҫ�޸ĵı����ͬ
                    index = i; // �����ͬ����¼��ǰ������������ѭ��
                    break;
                }
            }
            if (index != -1) { // ����ҵ���Ҫ�޸ĵ�����������Ҫ�޸ĵ���Ϣ�������������е�Ԫ�أ������浽�ļ��У�����ʾ�޸ĳɹ���Ϣ
                int newId, newDeptId; // �洢�µ�ְ����źͲ��ű��
                string newName; // �洢�µ�����
                string newduty; // �洢�µ�ְ��
                cout << "�������µ�ְ����ţ�" << endl;
                cin >> newId; // �����µ�ְ�����
                cout << "�������µ�������" << endl;
                cin >> newName; // �����µ�����
                cout << "�������µĲ��ű�ţ�" << endl;
                cin >> newDeptId; // �����µĲ��ű��
                cout << "�������µ�ְ��" << endl;
                cin >> newduty; // �����µ�ְ��
                Employee e(newId, newName, newDeptId, newduty); // ����һ���µ�ְ������
                employees[index] = e; // ���������е�ָ��Ԫ��
                saveFile(); // �������е����ݱ��浽�ļ���
                cout << "�ɹ��޸�ְ�����Ϊ" << id << "��ְ����" << endl;
            }
            else { // ���û���ҵ�Ҫ�޸ĵ���������ʾû���ҵ���ְ��
                cout << "û���ҵ�ְ�����Ϊ" << id << "��ְ����" << endl;
            }
        }

    }

    // ���ձ��������������ְ����ţ�������������������û�ָ��
    void sortEmployee() {
        if (employees.size() == 0) { 
            cout << "��ǰû�м�¼��������ӣ�" << endl;
        }
        else { 
            cout << "��ѡ������ʽ��1.���� 2.����" << endl;
            int choice; 
            cin >> choice; 
            if (choice == 1) { // ��������
                sort(employees.begin(), employees.end(), compareAsc); 
                cout << "����ְ�������������ɹ���" << endl;
            }
            else if (choice == 2) { // ��������
                sort(employees.begin(), employees.end(), compareDesc); 
                cout << "����ְ����Ž�������ɹ���" << endl;
            }
            else { 
                cout << "�����������������룡" << endl;
                return; // ��������
            }
            showEmployee(); // ��ʾ�����Ľ��
        }

    }

    // ��������ĵ�����������ļ��м�¼������ְ����Ϣ �����ǰ��Ҫ�ٴ�ȷ�ϣ���ֹ��ɾ��
    void clearFile() {
        cout << "ȷ����գ�1.ȷ�� 2.ȡ��" << endl;
        int choice; // �洢�û�ѡ���ȷ�ϻ�ȡ��
        cin >> choice; // �����û�ѡ���ȷ�ϻ�ȡ��
        if (choice == 1) { // ���ѡ��ȷ�ϣ����ļ��������������ļ����ݣ����ر��ļ�����������������������ʾ��ճɹ���Ϣ
            ofstream ofs(fileName, ios::trunc); // ���ļ��������ʹ�ýض�ģʽ
            ofs.close(); // �ر��ļ������
            employees.clear(); // �������
            cout << "�ɹ���������ĵ���" << endl;
        }
        else { 
            cout << "��ȡ����ղ�����" << endl;
        }

    }

    // �����ļ��������������е����ݱ��浽�ļ���
    void saveFile() {
        ofstream ofs(fileName, ios::out); // ���ļ��������ʹ�ø���ģʽ
        for (int i = 0; i < employees.size(); i++) { // ѭ�����������е�ÿ��Ԫ��
            Employee e = employees[i]; // ��ȡ��ǰְ������
            ofs << e.getId() << " "
                << e.getName() << " "
                << e.getDeptId() << " "
                << e.getduty() << endl; // ����ǰְ������Ϣд�뵽�ļ���
        }
        ofs.close(); // �ر��ļ������
    }

    // �ȽϺ�����������������
    static bool compareAsc(Employee e1, Employee e2) {
        return e1.getId() < e2.getId();
    }

    // �ȽϺ��������ڽ�������
    static bool compareDesc(Employee e1, Employee e2) {
        return e1.getId() > e2.getId(); 
    }
};

int main() {
    ManagementSystem ms("employee.txt"); // ����һ������ϵͳ���󣬲������ļ�����Ϊ����
    int choice; 
    while (true) { 
        ms.showMenu(); 
        cin >> choice; 
        switch (choice) { 
        case 0: // �˳�
            ms.exitSystem();
            break;
        case 1: // ����
            ms.addEmployee();
            break;
        case 2: // ��ʾ
            ms.showEmployee();
            break;
        case 3: // ɾ����ְְ��
            ms.deleteEmployee();
            break;
        case 4: // �޸�ְ����Ϣ
            ms.modifyEmployee();
            break;
        case 5: // ���ձ������
            ms.sortEmployee();
            break;
        case 6: // ��������ĵ�
            ms.clearFile();
            break;
        default: // ������Ч��ѡ���ʾ������Ϣ������������
            cout << "������������������" << endl;
        }
    }
        return 0;
    }
