import multiprocessing as mp 
import time 
import os 
def worker(i: int):
    print(f'{os.getpid()}: Working hard on {i}')
    time.sleep(2500)
    print(f'{os.getpid()}: Done with {i}!')

def main():
    print(f'Starting: {os.getpid()}')
    pool = mp.Pool(4)
    pool.map(worker, (0,1,2,3,))
    pool.close()
    pool.join()
    print('All done!')

if __name__ == '__main__':
    main()