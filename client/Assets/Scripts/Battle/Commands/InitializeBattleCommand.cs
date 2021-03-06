﻿using System;
using System.Linq;
using Zenject;
using UniRx;
using Type = TyphenApi.Type.Submarine;

namespace Submarine.Battle
{
    public class InitializeBattleCommand : Signal<InitializeBattleCommand>
    {
        public class Handler
        {
            [Inject]
            BattleService battleService;
            [Inject]
            BattleModel battleModel;
            [Inject]
            LobbyModel lobbyModel;

            public void Execute()
            {
                if (!lobbyModel.HasJoinedIntoRoom.Value)
                {
                    Logger.LogError("Not joined into a room");
                    return;
                }

                battleService.IsConnected.Where(v => v).Take(1).Subscribe(_ => OnBattleConnect());

                var room = lobbyModel.JoinedRoom.Value;
                var baseUri = new Uri(room.BattleServerBaseUri);
                var relativeUri = string.Format("rooms/{0}?room_key={1}", room.Id, room.RoomKey);
                battleService.Connect(new Uri(baseUri, relativeUri).ToString());
            }

            void OnBattleConnect()
            {
                battleModel.State.Value = BattleState.InPreparation;

                battleService.Api.OnRoomReceiveAsObservable().Subscribe(message =>
                {
                    var room = lobbyModel.JoinedRoom.Value;
                    room.Members = message.Members;
                    room.Bots = message.Bots;
                    lobbyModel.JoinedRoom.SetValueAndForceNotify(room);
                });

                battleService.Api.OnStartReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.StartedAt = CurrentMillis.FromMilliseconds(message.StartedAt);
                    battleModel.State.Value = BattleState.InBattle;
                });

                battleService.Api.OnFinishReceiveAsObservable().Take(1).Subscribe(message =>
                {
                    battleModel.Winner = lobbyModel.JoinedRoom.Value.Members.FirstOrDefault(m => m.Id == message.WinnerUserId);
                    battleModel.FinishedAt = CurrentMillis.FromMilliseconds(message.FinishedAt);
                    lobbyModel.JoinedRoom.Value = null;
                    battleModel.State.Value = BattleState.Finish;
                });

                battleService.Api.OnNowReceiveAsObservable().Subscribe(message =>
                {
                    battleModel.Now = CurrentMillis.FromMilliseconds(message.Time);
                });

                battleService.Api.OnActorReceiveAsObservable().Subscribe(message =>
                {
                    battleModel.ActorsById.Add(message.Id, message);
                });

                battleService.Api.OnDestructionReceiveAsObservable().Subscribe(message =>
                {
                    battleModel.ActorsById.Remove(message.ActorId);
                });

                battleService.Api.OnMovementReceiveAsObservable().Subscribe(message =>
                {
                    Type.Battle.Actor actor;
                    if (battleModel.ActorsById.TryGetValue(message.ActorId, out actor))
                    {
                        actor.UpdateValues(message);
                    }
                });

                battleService.Api.OnVisibilityReceiveAsObservable().Subscribe(message =>
                {
                    Type.Battle.Actor actor;
                    if (battleModel.ActorsById.TryGetValue(message.ActorId, out actor))
                    {
                        actor.UpdateValues(message);
                    }
                });

                battleService.Api.OnPingerReceiveAsObservable().Subscribe(message =>
                {
                    Type.Battle.Actor actor;
                    if (battleModel.ActorsById.TryGetValue(message.ActorId, out actor))
                    {
                        actor.UpdateValues(message);
                    }
                });

                battleService.Api.OnEquipmentReceiveAsObservable().Subscribe(message =>
                {
                    Type.Battle.Actor actor;
                    if (battleModel.ActorsById.TryGetValue(message.ActorId, out actor))
                    {
                        actor.UpdateValues(message);
                    }
                });
            }
        }
    }
}
