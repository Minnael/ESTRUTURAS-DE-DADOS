#include <iostream>
#include <queue>

int main(){
    //CRIANDO FILA
    std::queue <std::string> cartas;


    //INCLUINDO ELEMENTOS NA FILA
    cartas.push("REI DE COPAS");
    cartas.push("REI DE ESPADAS");
    cartas.push("REI DE OUROS");
    cartas.push("REI DE PAUS");


    //IMPRIMINDO TAMANHO DA FILA
    std::cout << "TAMANHO DA FILA: " << cartas.size() << std::endl;


    //RETIRANDO ELEMENTOS DA FILA (DA FRENTE -> "REI DE COPAS")
    cartas.pop();


    //RETORNANDO ELEMENTO DA FRENTE DA FILA
    std::cout << "CARTA DA FRENTE: " << cartas.front() << std::endl;


    //RETORNANDO ELEMENTO DE BAIXO DA FILA
    std::cout << "CARTA DE BAIXO: " << cartas.back() << std::endl;


    //VERIFICANDO SE A FILA ESTÁ VAZIA OU COM CARTAS
    if (cartas.empty()){
        std::cout << "FILA VAZIA" << std::endl;
    }
    else{
        std::cout << "FILA COM CARTAS" << std::endl;
    }


    //RETIRANDO TODOS OS ELEMENTOS DE UMA FILA
    while(!cartas.empty()){
        cartas.pop();
    }


    return 0;
}
