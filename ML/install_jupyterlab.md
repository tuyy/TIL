### linux에 jupyterlab 설치 및 remote 접속 가능하도록 설정 수정하기
0. python3을 설치한다.
    * ```sudo python3 -m pip install jupyterlab```
1. 최신 버전 https://github.com/jupyterlab/jupyterlab/releases 을 다운받는다.
2. 설치할 경로에 압축을 푼다. ```tar xvzf jupyterlab-0.34.10.tar.gz```
3. jupyterlab을 실행해 본다.
    * ```jupyter lab```
4. 원격 접속은 default로 제한하고 있어 설정을 수정해야한다. (default localhost만 허용)
    * ```jupyter notebook --generate-config``` // /home1/irteam/.jupyter/jupyter_notebook_config.py 가 생성됨
    * ```jupyter notebook password```  // password 설정
    * /home1/irteam/.jupyter/jupyter_notebook_config.py을 열고 
        * c.NotebookApp.ip = '*'
        * c.NotebookApp.password = '..' // /home1/irteam/.jupyter/jupyter_notebook_config.json의 password 부분을 붙여넣으면 된다. 
        * c.NotebookApp.open_browser = False
5. ```jupyter lab``` 다시 실행하여 원격 접속 여부를 확인한다.
