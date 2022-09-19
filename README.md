### Simple Go script to extract game names from XML gamelist files

#### It can be used with retro consoles based on EmulationStation as RetroPie. It reads the gamelist.xml files recursively, extract the game names and save it to a gamelist.txt file


#### Instructions:

1. The pre-compiled packages for RPI, Linux and Windows (not tested) can be found at `builds`

2. Usage: `getgamenames <path to gamelists files>` 

3. Example:

``` 
pi@retropie:~ $ ./getgamenames .emulationstation/gamelists/

 _____       _                _____              _
| __  | ___ | |_  ___  ___   |   __| ___  _____ |_| ___  ___
|    -|| -_||  _||  _|| . |  |  |  || .'||     || ||   || . |
|__|__||___||_|  |_|  |___|  |_____||__,||_|_|_||_||_|_||_  |
                                                        |___|

Starting game names extraction...
Done!!!
```

4. The `gamelist.txt` will be created in the same location


#### SonarQube Score

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aloula_get_game_names&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=aloula_get_game_names)

