#퀴즈
#기본 체력,코인
health = 5
coins = 0

print("게임 시작")

quiz = [("한국의 수도는?")]
answer = [("서울")}

#퀴즈 내는창
for health > 0: 
  print("\n문제:",quiz)
  user_answer = input("정답: ")

#카운트 다운
 for remaining in range(countdown_time, 0, -1):
            print(f"남은 시간: {remaining}초")
            time.sleep(1)
            if player.hp == 0:
                break
        print(' ' * 20, end='\r')#카운트 삭제  
          
#퀴즈 정답
if user_answer == answer:
        coins += 1
        print("정답! 코인 1개 획득!")
    #퀴즈땡
          else:
          health -= 1
          print(f"틀렸어요! 체력 1 감소! (남은 체력: {health})")
if health == 0:#죽음
  print("game over")
  print("얻은 코인 갯수: {coins}")
          #카운트 정답이면 초기화
          if user_input.strip() == answer:
            player.reward()
            return = countdown_time
        else:
            player.penalty()
