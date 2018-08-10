## scikit-learn 개발 환경 통합 패키지 
* Anaconda (NumPy, SciPy, matplotlib, pandas, IPython, 주피터 노트북, scikit-learn 모두 포함)
* macOS, 윈도우, 리눅스 모두 지원
* www.continuum.io/anaconda-overview

### 또는 python 3.6이 설치되었다면,

```
pip install numpy scipy matplotlib ipython scikit-learn pandas pillow
```

## scikit-learn에서 사용하는 필수 라이브러리 소개

### NumPy
* pip install numpy
* 파이썬으로 과학 계산을 할때 꼭 필요한 패키지
* scikit-learn에서 NumPy배열은 기본 데이터 구조이다. (핵심인 다차원 배열이며 ndarray 클래스로 정의)

```
import numpy as np

x = np.array([[1, 2, 3], [4, 5, 6]])
print('x:\n{}'.format(x))
```

### SciPy
* 과학 계산용 함수를 모아놓은 파이썬 패키지
* 고성능 선형 대수, 함수 최적화, 신호 처리, 특수한 수학 함수와 통계 분포 등 함수 제공

```
from scipy import sparse

eye = np.eye(4)
print('NumPy 배열\n{}\n'.format(eye))

sparse_matrix = sparse.csr_matrix(eye)
print('sparse_matrix: \n{}'.format(sparse_matrix))
```

### matplotlib
* 파이썬의 대표적인 과학 계산용 그래프 라이브러리
* 선 그래프, 히스토그램, 산점도 등을 지원하며 출판에 쓸 수 있을 만큼의 고품질 그래프를 그려준다.
* 데이터와 분석 결과를 다양한 관점에서 시각화해보면 매우 중요한 통찰을 얻을 수 있다.

```
import matplotlib.pyplot as plt

x = np.linspace(-10, 10, 100)

y = np.sin(x)

plt.plot(x, y, marker='x')
Out[17]: [<matplotlib.lines.Line2D at 0x108eb1eb8>]
￼# sin 그래프가 표시된다.
```

### pandas
* 데이터 처리와 분석을 위한 파이썬 라이브러리
* R의 data.frame을 본떠서 설계한 DataFrame이라는 데이터구조를 기반으로 만들어졌다.
* 간단히 말하면 pandas의 DataFrame은 엑셀의 스프레드시트와 비슷한 테이블 형태라고 할 수 있다.
* SQL처럼 테이블에 쿼리나 조인을 수행할 수 있다.
* CSV, 엑셀파일을 쉽게 읽을 수 있다.

```
from IPython.display import display
import pandas as pd

data = {'name' : ['tu', 'yy', 'jay'],
        'age'  : [11, 12, 13],
        'location' : ['kr', 'kr', 'jp']}

data_pandas = pd.DataFrame(data)
display(data_pandas)

print()

display(data_pandas[data_pandas.age > 12])
```
