/*
Code By:
Diogo "Joe" Delazare Brandão - 2022

This application was made possible by:
The Go proggraming language (Golang) by Google,
The PokeAPI API by ...
*/

package main

import (
	"Joe/sheet-hole/pkg/pokemon/PTA1"
	"fmt"
	"log"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache string //map[string]*template.Template
}

func main() {
	// infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// app := &application{
	// 	infoLog:  infoLogger,
	// 	errorLog: errorLogger,
	// }

	//////////////////////////////////////////////////////////////////////////////////////////////

	testExpertise, _ := PTA1.RegisterExpertise("test expertise", "ATK", "Test expertise")

	base1, _ := PTA1.RegisterTrainerTalent("Pokéagenda", false, "Nenhum", "À vontade", "Um pokemon ainda não identificado por sua pokeagenda a até 10 metros", "apontando sua pokéagenda a um pokémon não identificado, ela o identifica. Uma vez que um pokémon possui variações de aparência – desde forma, coloração, pelugem, etc. – mesmo dentro de uma mesma espécie, uma pokéagenda demora um pouco para fazer a análise visual. O tempo necessário é de 12 segundos, ou duas rodadas de combate, para que ela lhe dê as informações. Pokéagendas são aparelhos personalizados, e constituem a identificação oficial de um alguém, portanto são licenciadas para serem entregues a pessoas a partir de 10	anos, e se espera que estejam sempre portadas por alguém para que este indivíduo não tenha problemas com autoridades. Elas são produzidas pela Liga Pokémon e pelas autoridades governamentais e informam a todos que aquele indivíduo possui autorização legal para passear por aí, é quem ele diz ser e porta pokémons. Elas são atualizadas por pesquisadores profissionais e contribuintes leigos. Todas as pokéagendas são a prova d’água, e extremamente	resistentes a calor, ácido e muitas outras formas de dano. Costuma ser do melhor interesse de alguém manter uma	pokéagenda íntegra, e as autoridades terão que se envolver se uma pokéagenda for destruída, perdida ou roubada.", false, false, true, false, false, true)
	base2, _ := PTA1.RegisterTrainerTalent("Pokéagenda2", false, "Nenhum", "À vontade", "Um pokemon ainda não identificado por sua pokeagenda a até 10 metros", "apontando sua pokéagenda a um pokémon não identificado, ela o identifica. Uma vez que um pokémon possui variações de aparência – desde forma, coloração, pelugem, etc. – mesmo dentro de uma mesma espécie, uma pokéagenda demora um pouco para fazer a análise visual. O tempo necessário é de 12 segundos, ou duas rodadas de combate, para que ela lhe dê as informações. Pokéagendas são aparelhos personalizados, e constituem a identificação oficial de um alguém, portanto são licenciadas para serem entregues a pessoas a partir de 10	anos, e se espera que estejam sempre portadas por alguém para que este indivíduo não tenha problemas com autoridades. Elas são produzidas pela Liga Pokémon e pelas autoridades governamentais e informam a todos que aquele indivíduo possui autorização legal para passear por aí, é quem ele diz ser e porta pokémons. Elas são atualizadas por pesquisadores profissionais e contribuintes leigos. Todas as pokéagendas são a prova d’água, e extremamente	resistentes a calor, ácido e muitas outras formas de dano. Costuma ser do melhor interesse de alguém manter uma	pokéagenda íntegra, e as autoridades terão que se envolver se uma pokéagenda for destruída, perdida ou roubada.", false, false, true, false, false, true)

	class, _ := PTA1.RegisterTrainerClass("testClass", "Just a test class", "", [2]*PTA1.TrainerTalent{base1, base2}, []*PTA1.TrainerTalent{base1}, []*PTA1.Expertise{testExpertise}, "test Requirements")

	fmt.Printf("\n\n%+v\n\n", class)
}
