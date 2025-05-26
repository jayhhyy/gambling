#창 에서 버튼 클릭
import tkinter as tk
from tkinter import messagebox


#퀴즈
# 처음값
health = 5
coins = 0
score = 0
countdown_time = 5

print("게임 시작!")

# 퀴즈 리스트 더 추가하기$$$$$$$$$
quiz_list = [("한국의 수도는?", "서울"),
             ("2 + 2는?", "4"),
             ("가장 높은 산은?","에베레스트산"),
             ("가장 큰 바다는?", "태평양")]

# 퀴즈 내기
while health > 0:
    import random
    question, answer = random.choice(quiz_list)
    
    # 카운트 다운
    for remaining in range(countdown_time, 0, -1):
        print("남은 시간: {remaining}초", end='\r')
        time.sleep(1)
        if health <= 0:
            break
    print(' ' * 30, end='\r')  # 카운트다운 지우기

    if health <= 0:
        break

    print(f"\n문제: {question}")
    user_answer = input("정답: ")

    if user_answer.strip() == answer:
        coins += 1
        score += 1
        countdown_time = 5  # 정답시 카운트다운 다시
        print("정답! 코인 1개 획득!")
    else:
        health -= 1
        score += 1
        print("틀렸어요! 체력 1 감소! (남은 체력: {health})")

#마지막에 푼 퀴즈수와 얻은 코인수
print("\nGame Over")
print("얻은 코인 개수: {coins}")
print("푼 퀴즈 개수: {score}")



