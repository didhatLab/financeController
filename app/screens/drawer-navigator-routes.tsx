import {createStackNavigator} from '@react-navigation/stack';
import {createDrawerNavigator} from '@react-navigation/drawer';

import {RootStackParamList} from "../types/screen-types";
import {NativeStackScreenProps} from "@react-navigation/native-stack";
import {useEffect, useState} from "react";
import AsyncStorage from "@react-native-async-storage/async-storage";
import {HomeScreen} from "./menu/spend/home";
import {DrawerParams} from "./menu/types";
import {AddNewSpendScreen} from "./menu/spend/add-spending";
import {HomeRoutes} from "./menu/spend/route";
import {GroupRoutes} from "./group";

type Props = NativeStackScreenProps<RootStackParamList, 'DrawNavigationRoutes'>

const Drawer = createDrawerNavigator<DrawerParams>()
const Stack = createStackNavigator()


export const DrawerNavigatorRoutes = (props: Props) => {
    return (
        <Drawer.Navigator>
            <Drawer.Screen name={"Home"} component={HomeRoutes}
                           options={{drawerLabel: 'Home', title: 'My Spends', headerTitle: 'My Spends'}}
            />
            <Drawer.Screen name={"Groups"} component={GroupRoutes}/>
        </Drawer.Navigator>
    )
}
