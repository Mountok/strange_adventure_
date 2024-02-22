package frames

import (
	"fmt"
	"strangeadventure/data/colors"
	"strangeadventure/models/player"
)


var GetScreen = map[string]func()string{
	
	"menu_1": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"# ———————————————————|\n",
			"# Strange Adventure\n",
			"# ———————————————————|\n",
			"# [",colors.ColorRed(),">",colors.ColorReset(),"]",colors.ColorGreen(),"	Играть",colors.ColorReset(),"\n",
			"# [ ] Персонаж\n",
			"# [ ] Помошь\n",
			"# [ ] Выйти\n",
			"# ———————————————————|\n",
			"#####################################################################################################\n",			
		)
	},
	"menu_2": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"# ———————————————————|\n",
			"# Strange Adventure\n",
			"# ———————————————————|\n",
			"# [ ] Играть\n",
			"# [",colors.ColorRed(),">",colors.ColorReset(),"]",colors.ColorGreen(),"	Персонаж",colors.ColorReset(),"\n",
			"# [ ] Помошь\n",
			"# [ ] Выйти\n",
			"# ———————————————————|\n",
			"#####################################################################################################\n",		
		)
	},
	"menu_3": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"# ———————————————————|\n",
			"# Strange Adventure\n",
			"# ———————————————————|\n",
			"# [ ] Играть\n",
			"# [ ] Персонаж\n",
			"# [",colors.ColorRed(),">",colors.ColorReset(),"]",colors.ColorGreen(),"	Помошь",colors.ColorReset(),"\n",
			"# [ ] Выйти\n",
			"# ———————————————————|\n",
			"#####################################################################################################\n",		
		)
	},
	"menu_4": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"# ———————————————————|\n",
			"# Strange Adventure\n",
			"# ———————————————————|\n",
			"# [ ] Играть\n",
			"# [ ] Персонаж\n",
			"# [ ] Помошь\n",
			"# [",colors.ColorRed(),">",colors.ColorReset(),"]",colors.ColorRed(),"	Выйти",colors.ColorReset(),"\n",
			"# ———————————————————|\n",
			"#####################################################################################################\n",	
		)
	},
	"inter_name": func() string {
		return fmt.Sprint(
			"################################################################\n",
			"# ———————————————————|\n",
			"#",colors.ColorRed()," Создание персонажа\n",colors.ColorReset(),
			"# ———————————————————|\n",
			"#",colors.ColorGreen()," Имя: ",colors.ColorReset(),
		)
	},
	"world_0_0": func () string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"#                                                                                                   #\n",
			"#  |                 |                                                                              #\n",
			"#  |  [лес мертвых]  |                                                                              #\n",
			"#  |                 |                                                                              #\n",
			"#                                [][][][][][-----][][][][][]                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []  	  [СТОЛИЦА]       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                [][][][][][]---[][][][][][]                                        #\n",
			"#                                                                                                   #\n",
			"#  |               |                                                                                #\n",
			"#  |   [мародеры]  |                                                                                #\n",
			"#  |               |                                                                                #\n",
			"#                                                                                                   #\n",
			"#####################################################################################################\n",
			"# ",colors.ColorCyan(),"Выбор локации стрелки ВНИЗ/ВВЕРХ.",colors.ColorReset(),
			"                                                    ",
			colors.ColorRed(),"Выход - Enter",colors.ColorReset(),"#\n",			
			"#####################################################################################################",
		)
	},
	"world_0_1": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"#                                                                                                   #\n",
			"#  |                 |                                                                              #\n",
			"#  |  ",
			colors.ColorRed(),"[лес мертвых]",colors.ColorReset(),
			"  |                                                                              #\n",
			"#  |                 |                                                                              #\n",
			"#                                [][][][][][-----][][][][][]                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []  	  [СТОЛИЦА]       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                [][][][][][]---[][][][][][]                                        #\n",
			"#                                                                                                   #\n",
			"#  |               |                                                                                #\n",
			"#  |   [мародеры]  |                                                                                #\n",
			"#  |               |                                                                                #\n",
			"#                                                                                                   #\n",
			"#####################################################################################################\n",
			"# ",colors.ColorCyan(),"Выбор локации стрелки ВНИЗ/ВВЕРХ.",colors.ColorReset(),
			"                                                    ",
			colors.ColorRed(),"Выход - Enter",colors.ColorReset(),"#\n",			
			"#####################################################################################################",			
		)
	},
	"world_0_2": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"#                                                                                                   #\n",
			"#  |                 |                                                                              #\n",
			"#  |  [лес мертвых]  |                                                                              #\n",
			"#  |                 |                                                                              #\n",
			"#                                [][][][][][-----][][][][][]                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []  	  ",
			colors.ColorGreen(),"[СТОЛИЦА]",colors.ColorReset(),
			"       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                [][][][][][]---[][][][][][]                                        #\n",
			"#                                                                                                   #\n",
			"#  |               |                                                                                #\n",
			"#  |   [мародеры]  |                                                                                #\n",
			"#  |               |                                                                                #\n",
			"#                                                                                                   #\n",
			"#####################################################################################################\n",
			"# ",colors.ColorCyan(),"Выбор локации стрелки ВНИЗ/ВВЕРХ.",colors.ColorReset(),
			"                                                    ",
			colors.ColorRed(),"Выход - Enter",colors.ColorReset(),"#\n",			
			"#####################################################################################################",			
		)
	},
	"world_0_3": func() string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"#                                                                                                   #\n",
			"#  |                 |                                                                              #\n",
			"#  |  [лес мертвых]  |                                                                              #\n",
			"#  |                 |                                                                              #\n",
			"#                                [][][][][][-----][][][][][]                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []  	  [СТОЛИЦА]       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                []                       []                                        #\n",
			"#                                [][][][][][]---[][][][][][]                                        #\n",
			"#                                                                                                   #\n",
			"#  |               |                                                                                #\n",
			"#  |   ",
			colors.ColorRed(),"[мародеры]",colors.ColorReset(),
			"  |                                                                                #\n",
			"#  |               |                                                                                #\n",
			"#                                                                                                   #\n",
			"#####################################################################################################\n",
			"# ",colors.ColorCyan(),"Выбор локации стрелки ВНИЗ/ВВЕРХ.",colors.ColorReset(),
			"                                                    ",
			colors.ColorRed(),"Выход - Enter",colors.ColorReset(),"#\n",			
			"#####################################################################################################",		
		)
	},
	
}

var GetScreenWithPlayerInfo = map[string]func(p player.Player)string{
	"profile": func(p player.Player) string {
		return fmt.Sprint(
			"#####################################################################################################\n",
			"# ———————————————————|\n",
			"# Ваш персонаж \n",
			"# ———————————————————|\n",
fmt.Sprintf("# [Имя] %s\n",p.GetInfo().Name),
fmt.Sprintf("# [Уровень] %d\n",p.GetInfo().Level),
fmt.Sprintf("# [Ранг] %s\n",p.GetInfo().Rank),
fmt.Sprintf("# [Жизни] %d\n",p.GetInfo().Health),
fmt.Sprintf("# [Мир] %d\n",p.GetInfo().World),
			"# ———————————————————|\n",
			"# Esc - Выйти\n",
			"# ———————————————————|\n",
			"#####################################################################################################\n",	
		)
	},
}

