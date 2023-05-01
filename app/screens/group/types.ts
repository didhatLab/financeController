import React from "react";
import {createNativeStackNavigator, NativeStackScreenProps} from "@react-navigation/native-stack";
import {DrawerParams} from "../menu/types";
import {Spending} from "../../types/data-types";


export type GroupStackParams = {
    GroupList: {}
    GroupScreen: {group: SpendGroup}
    EditGroupSpend: {spend: Spending}
}

export type GroupPropsForRoute = NativeStackScreenProps<GroupStackParams, 'GroupList'>;


export interface GroupMember {
    UserId: number
    Username: string
    IsAdmin: boolean
}

export interface SpendGroup {
    Id: number
    Name: string
    Description: string
    Members: GroupMember[]
}

export type GroupContextType = {
    groups: SpendGroup[]
    setGroups: React.Dispatch<React.SetStateAction<SpendGroup[]>>
}


export type GroupRouteProps = NativeStackScreenProps<DrawerParams, 'Groups'>
export const GroupStack = createNativeStackNavigator<GroupStackParams>()