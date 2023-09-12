#include <iostream>
#include <stack>

int main(){
    //CRIANDO PILHA
    std::stack <std::string> cartas;


    //INCLUINDO ELEMENTOS NA PILHA
    cartas.push("REI DE COPAS");
    cartas.push("REI DE ESPADAS");
    cartas.push("REI DE OUROS");
    cartas.push("REI DE PAUS");


    //IMPRIMINDO TAMANHO DA PILHA
    std::cout << "TAMANHO DA PILHA: " << cartas.size() << std::endl;


    //RETIRANDO ELEMENTOS DA PILHA (DO TOPO -> "REI DE PAUS")
    cartas.pop();


    //RETORNANDO ELEMENTO DO TOPO DA PILHA
    std::cout << "CARTA DO TOPO: " << cartas.top() << std::endl;


    //VERIFICANDO SE A PILHA ESTÁ VAZIA OU COM CARTAS
    if (cartas.empty()){
        std::cout << "PILHA VAZIA" << std::endl;
    }
    else{
        std::cout << "PILHA COM CARTAS" << std::endl;
    }


    //RETIRANDO TODOS OS ELEMENTOS DE UMA PILHA
    while(!cartas.empty()){
        cartas.pop();
    }

    std::cout << cartas.empty() << std::endl;

    return 0;
}
