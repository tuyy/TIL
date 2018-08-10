## 붓꽃의 품종 분류 예제

### 개요
* 붓꽃의 품종을 분류하는 모델을 만든다.
* 붓꽃의 꽃잎과 꽃받침의 폭과 길이를 센티미터 단위로 측정한 값들로 품종을 분류한다.
* 품종은 setosa, versicolor, virginica 종으로 분류한다.

### 머신러닝 기본 용어
* 붓꽃의 품종을 정확하게 분류한 데이터를 가지므로 **지도 학습**에 속한다.
* 또한 이 문제의 경우엔 몇 가지 선택사항(품종) 중 하나를 선택하는 문제이므로 **분류 문제(classification)** 에 속한다.
* 출력될 수 있는 값(결과)을 **클래스(class)** 라 한다. -> 이 문제에선 setosa, versicolor, virginica
* 데이터 포인트(one row) 하나(붓꽃 하나)에 대한 기대 출력은 꽃의 품종이 된다. 이런 특정 데이터 포인트에 대한 출력, 즉 품종을 **레이블(label)** 이라 한다.
* 머신러닝에서 각 붓꽃의 데이터들을 **샘플** 이라 하고 속성은 **특성(feature)** 이라 한다.

### 붓꽃 dataset에 대한 확인
```
# -*- coding: utf-8 -*-

from sklearn.datasets import load_iris

iris_dataset = load_iris()

#print("iris_dataset의 key: \n{}".format(iris_dataset.keys()))

# iris_dataset에 대한 설명
#print(iris_dataset['DESCR'][:193] + '\n...')

# 클래스
#print('target_names: ')
#print(iris_dataset['target_names'])

# 속성 종류
#print('features names:\n{}'.format(iris_dataset['feature_names']))

# 샘플
#print('data의 타입: {}'.format(type(iris_dataset['data'])))
#print('data의 크기: {}'.format(iris_dataset['data'].shape))
#print('data의 처음 다섯 행:\n{}'.format(iris_dataset['data'][:5]))

# 샘플에 대한 레이블
#print("타깃:\n{}".format(iris_dataset['target']))
```

### 성과 측정
* 우리가 만든 모델의 예측을 신뢰할 수 있으려면 성과 측정해야한다.
* 불행히 모델을 만들 때 쓴 데이터는 평가 목적으로 사용할 수 없다. 전부 기억할 수 있기 때문에 정확히 맞출 수도 있기 때문이다.
* 이렇게 데이터를 기억한다는 것은 모델을 잘 **일반화** 하지 않았다는 뜻이다. 즉, 새로운 데이터에 대해서는 잘 동작하지 않는다.
* 모델의 성능을 측정하려면 레이블을 알고 있는 새 데이터를 모델에 적용해봐야한다.
* 모델을 만들 때 사용하는 데이터를 훈련 데이터 혹은 **훈련 세트(training set)** 이라 한다.
* 모델이 얼마나 잘 작동하는지 측정하는데 사용하는 데이터를 테스트 데이터, **테스트 세트(test set)** , **혹은 홀드아웃 세트(hold-out set)** 라 한다.
* 일반적으로 전체 데이터의 훈련은 75%, 테스트는 25%를 사용하는 것은 좋은 선택이다.

### train, test dataset 분리하기
* scikit-learn에서는 데이터는 대문자X로 표시하고 레이블은 소문자 y로 표시한다.

```
# train, test dataset 나누기 (train 75 : test 25)
from sklearn.model_selection import train_test_split
X_train, X_test, y_train, y_test = train_test_split(
        iris_dataset['data'], iris_dataset['target'], random_state=0)

print('X_train size: {}'.format(X_train.shape))
print('y_train size: {}'.format(y_train.shape))
print('X_test size: {}'.format(X_test.shape))
print('y_test size: {}'.format(y_test.shape))
```

### 가장 먼저 할 일: 데이터 살펴보기
* 일단 모델을 만들기 전에 머신러닝이 없이도 풀 수 있는 문제는 아닌지, 혹은 필요한 정보가 누락되지는 않았는지 데이터를 조사해보는 것이 필요하다.
* 또는 비정상 적인 데이터가 있는지 확인해야한다. 이를 위해 데이터 시각화를 해보면 좋다.
* 산점도(scatter plot)은 데이터에서 한 특성을 x 축에 놓고 다른 하나는 y 축에 놓아 각 데이터 포인트를 하나의 점으로 나타내는 그래프이다.
* 주의할 점은 산점도 행렬은 한 그래프에 모든 특성의 관계가 나타나는 것이 아니기 떄문에 각가의 나누어진 산점도 그래프에는 드러나지 않는 중요한 성질이 있을 수 있다.

```
iris_dataframe = pd.DataFrame(X_train, columns=iris_dataset.feature_names)
pd.plotting.scatter_matrix(iris_dataframe, c=y_train, figsize=(15, 15),
                           marker='o', hist_kwds={'bins': 20}, s=60, alpha=.8,
                           cmap=mglearn.cm3)
```

### 첫 번째 머신러닝 모델: k-최근접 이웃 알고리즘
* k-최근접 이웃(k-Nearest Neighbors, k-NN) 분류기를 이용하면 단순히 훈련 데이터를 저장하여 모델을 만들 수 있다.
* 새로운 데이터 포인트에 대한 예측이 필요하면 알고리즘은 새 데이터 포인트에서 가장 가까운 훈련 데이터 포인트를 찾아 새 데이터 포인트의 레이블로 지정한다.
* 여기서 k는 가장 가까운 이웃 '하나'가 아니라 훈련 데이터에서 새로운 데이터 포인트에 가장 가까운 'k개'의 이웃을 찾는다는 뜻이다.
* k-NN 알고리즘에서 가장 중요한 매개변수는 이웃의 개수이다.

```
from sklearn.model_selection import train_test_split
X_train, X_test, y_train, y_test = train_test_split(
        iris_dataset['data'], iris_dataset['target'], random_state=0)

from sklearn.neighbors import KNeighborsClassifier
knn = KNeighborsClassifier(n_neighbors=1)

knn.fit(X_train, y_train)

X_new = np.array([[5, 2.9, 1, 0.2]]) # 정확한 레이블을 모르는 샘플 데이터
print("X_new.shape: {}".format(X_new.shape))

prediction = knn.predict(X_new) # 예측!
print("예측: {}".format(prediction))
print("예측한 타깃의 이름: {}".format(iris_dataset['target_names'][prediction]))
```

### 모델 평가하기
* 평가를 위해 test set 을 사용하면 된다.
* 테스트 데이터에 있는 붓꽃의 품종을 예측하고 실제 레이블과 비교하여 정확도를 계산하면 모델의 성능을 평가 할 수 있다.

```
y_pred = knn.predict(X_test)
print("테스트 세트에 대한 예측값:\n{}".format(y_pred))

print("정확도: {:.2f}".format(np.mean(y_pred == y_test))) # 0.97, 즉 97%의 정확도를 가진다.
```

### 최종! 훈련과 평과 과정을 담은 전체 코드

```
from sklearn.model_selection import train_test_split
from sklearn.neighbors import KNeighborsClassifier

X_train, X_test, y_train, y_test = train_test_split(
        iris_dataset['data'], iris_dataset['target'], random_state=0) # 훈련 데이터와 테스트 데이터 분리

knn = KNeighborsClassifier(n_neighbors=1) # 알고리즘 선택

knn.fit(X_train, y_train) # 훈련..

print("테스트 세트 정확도: {:.2f}".format(knn.score(X_test, y_test))) # 평가..
```

### 참고자료
* introduction to Machine Learning with Python - 안드레아스 뮐러, 세라 가이도 
