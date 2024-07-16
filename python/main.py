import requests

if __name__ == '__main__':
    port = 3000
    while True:
        user_input = input('Enter a number: ')
        try:
            res = requests.get(f'http://localhost:{port}/br/{user_input}')
            res_json = res.json()
        except:
            print('Start Server & Try Again')
            quit()

        print(f'Your Input: {res_json["val"]}')
        if res_json["err"]:
            print('Invalid Input')
            continue
        try:
            print(f'Output: {res_json["txt"]}')
        except KeyError:
            print('Invalid Input')
            continue
        
