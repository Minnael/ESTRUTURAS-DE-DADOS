#include <iostream>
#include <list>


int main(){
    std::list <int> aula;
    int tam = 10;

    for(int i=0; i<10; i++){
        aula.push_front(i);
    }

    std::cout<<"tamanho da lista: " << aula.size() << std::endl;

    tam = aula.size();
    for(int i=0; i<tam; i++){
        std::cout << aula.front() << std::endl;
        aula.pop_front();
    }

    return 0;
}
