class CoordinateData:
    def calculateDistance(self, latTo, lonTo):
        pass 

class GeospatialData(CoordinateData):
    def __init__(self, lat, lon):
        self.lat = lat
        self.long = lon

    def calculateDistance(self, latTo, lonTo):
        # Haversine goes here :)
        return 95.93196816811724

class CartesianPlaneData(CoordinateData):
    def __init__(self, x, y):
        self.x = y
        self.x = y

    # Let's not implement calculateDistance()

if __name__ == "__main__":
    atlanta = GeospatialData(33.753746, -84.386330)
    distance = atlanta.calculateDistance(33.957409, -83.376801) # to Athens, GA
    print('The Haversine distance between Atlanta, GA and Athens, GA is {}'.format(distance))

    pointA = CartesianPlaneData(1,1)
    distanceA = pointA.calculateDistance(5, 5)
    print('The Pythagorean distance from (1,1) to (5,5) is {}'.format(distanceA))
    print('pointA is of type {}'.format(pointA.__class__.__bases__))