import 'react-native-gesture-handler';
import React from 'react';
import { StyleSheet, Text, View, FlatList, Image ,

 Pressable 

} from 'react-native';

import Home from './screens/Home';
import ColorPalette from './screens/ColorPalette';
import AddNewPaletteModal from './screens/AddNewPaletteModal';


import { NavigationContainer, useNavigation} from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

import { WebView } from 'react-native-webview';


import { styled } from 'nativewind';


function Button(props) {
  const { onPress, title = 'Save' } = props;
  return (
    <Pressable style={styles.button} onPress={onPress}>
      <Text style={styles.text}>{title}</Text>
    </Pressable>
  );
}

const StyledView = styled(View)
const StyledText = styled(Text)



const styles = StyleSheet.create({
  button: {
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 12,
    paddingHorizontal: 32,
    borderRadius: 4,
    elevation: 3,
    backgroundColor: 'black',
  },
  text: {
    fontSize: 16,
    lineHeight: 21,
    fontWeight: 'bold',
    letterSpacing: 0.25,
    color: 'white',
  },
  chrome : {
    height: 60, // Reduced height for a more compact view
    paddingTop: 10, // Adjust padding to center content vertically within the smaller height
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
    //height: '100%'
  },
  gridContainer: {
    flex: 1,
    margin: 2,
  },
  itemStyle: {
    flex: 1,
    margin: 5,
    minWidth: 150,
    maxWidth: 150,
    height: 100,
    maxHeight: 100,
    // backgroundColor: '#CCC',
  },
});


const ComponentNine= () => {
  //
  return <WebView source={{ uri: 'https://transmit.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentOne</Text>
     </View>
  </>)
}

const ComponentTen= () => {
  return <WebView source={{ uri: 'https://pocket.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentOne</Text>
     </View>
  </>)
}

const ComponentEleven = () => {
  return <WebView source={{ uri: 'https://syntax.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentOne</Text>
     </View>
  </>)
}


const ComponentTwelve = () => {
  return <WebView source={{ uri: 'https://keynote.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentOne</Text>
     </View>
  </>)
}

const ComponentTwo = () => {
  return <WebView source={{ uri: 'https://primer.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentTwo</Text>
     </View>
  </>)
}

const ComponentThree = () => {

  return <WebView source={{ uri: 'https://commit.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentThree</Text>
     </View>
  </>)
}

const ComponentFour = () => {
  return <WebView source={{ uri: 'https://protocol.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentFour</Text>
      
     </View>
  </>)
}

const ComponentFive = () => {

  return <WebView source={{ uri: 'https://salient.tailwindui.com/' }} />;

  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentFive</Text>
     </View>
  </>)
}

const ComponentSix = () => {
  return (<>
    <View style={styles.container}>
      <Text>Hello ComponentSix</Text>
      <Image
        source={require('./assets/6.webp')}
        style={{width: "100%", height: "100%"}}

      />
     </View>
  </>)
}

const ComponentSeven = () => {
  return (<>
    <View style={styles.container}>
      <Text>Hello Shopping Mall</Text>
      <Image
        source={require('./assets/7.jpeg')}
        style={{width: "100%", height: "100%"}}
      />
     </View>
  </>)
}

const ComponentEight = () => {
  return (<>
    <View style={styles.container}>
    <Text>Hello ATAK-WINTAKx</Text>

      <Image
        source={require('./assets/8.webp')}
        style={{width: "95%", height: "90%"}}

      />
     </View>
  </>)
}


const ComponentOne = () => {
  return (<>
    <View style={styles.container}>
    <Text className="font-bold text-xl">Hello Swarmbotics.AI</Text>
    <Text>
    <Text style={{marginBottom: 10}}>
      <Text style={{fontWeight: 'bold', fontSize: 18, marginTop: 20}}>
        Hi, My name is Adnan.{'\n\n'}
      </Text>
      <Text style={{fontSize: 16}}>
        I've been a software developer for 12 years.{'\n'}
        I love UI/UX Design, geospatial mapping, Hardware, and MLOps for GPU Inference.{'\n\n'}
        On the agenda, I am presenting 3 items today.{'\n\n'}
      </Text>
      <Text style={{fontSize: 16, fontWeight: 'bold'}}>
        1. Simple Android App I started on Feb 24.{'\n\n'}
      </Text>
      <Text style={{fontSize: 16}}>
        It's designed to assist with co-operative robotics control for medical bots, house cleaning bots and more.{'\n\n'}
      </Text>
      <Text style={{fontSize: 16, fontWeight: 'bold'}}>
        2. Figma / Dribble with 10 examples of designs I've built over the years.{'\n\n'}
      </Text>
      <Text style={{fontSize: 16}}>
        3. Some WebGL Demos @ adnanwahab.com{'\n\n'}
      </Text>
      <Text style={{fontSize: 16}}>
        1 & 2 will take 15 min each. 3 will take 5 min.{'\n\n'}
      </Text>
      <Text style={{color: 'blue', fontSize: 16}}
            onPress={() => Linking.openURL('https://www.figma.com/file/TV2fZrBFw5tSkKMEVQ29hR/SwarmBotics-Presentation---Feb-26?type=design&node-id=0-1&mode=design&t=Uyyif5IylYrppsMZ-0')}>
        Click here to view the Figma presentation
      </Text>
    </Text>
    </Text>
     </View>
  </>)
}

const Stack = createStackNavigator();


function HomeScreen({ navigation }) {
  let numWords = 'One two three four five six seven eight nine ten eleven twelve'.split(' ')
  numWords = numWords.map(word => word.charAt(0).toUpperCase() + word.slice(1));
  
  // .map(_ => {
  //   let wtf = _.slice()
  //   wtf[0] = wtf[0].toUpperCase()

  //   return wtf
  // })

  let content = numWords.map((word, index) => {
    return {
      key: `${index}`,
      title: `Go to Screen ${index + 1}`,
      navigateTo: word,
    }
  });

  return (
    <FlatList
      data={content}
      renderItem={({ item }) => (
        <View style={styles.itemStyle}>
          <Button
            title={item.title}
            onPress={() => navigation.navigate(item.navigateTo)}
          />
        </View>
      )}
      numColumns={2}
      keyExtractor={(item) => item.key}
      contentContainerStyle={styles.gridContainer}
    />
  );
}


const TableOfContents = () => {
  return (
    <>
    {/* <View style={styles.chrome}>
      <Text>Hello Swarmbotics.AI</Text>
     </View> */}
    <View className="flex-1 justify-center items-center pt-5">
    <Text className="font-bold text-xl pt-5"></Text>
    <Text className="font-bold text-xl pt-5"></Text>

      <Text className="font-bold text-xl pt-5 pl-5">                                 Hello, Swarmbotics.ai!</Text>
    </View>
        <Stack.Navigator>
          <Stack.Screen name="Home" component={HomeScreen} />

          <Stack.Screen name="One" component={ComponentOne} />


          <Stack.Screen name="Two" component={ComponentTwo} />
          <Stack.Screen name="Three" component={ComponentThree} />
          <Stack.Screen name="Four" component={ComponentFour} />

          <Stack.Screen name="Five" component={ComponentFive} />
          <Stack.Screen name="Six" component={ComponentSix} />
          <Stack.Screen name="Seven" component={ComponentSeven} />
          <Stack.Screen name="Eight" component={ComponentEight} />
          <Stack.Screen name="Nine" component={ComponentNine} />
          <Stack.Screen name="Ten" component={ComponentTen} />
          <Stack.Screen name="Eleven" component={ComponentEleven} />
          <Stack.Screen name="Twelve" component={ComponentTwelve} />

      </Stack.Navigator>
    </>
  )
}

const App = () => {
  return (
    <NavigationContainer>
      <TableOfContents />
    </NavigationContainer>
  );
}

export default App;
