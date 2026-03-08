import { Injectable } from '@angular/core';
import { Entity } from '../../core/models/entity';
import { Observable, of} from 'rxjs';

export interface Character extends Entity {}

@Injectable({
  providedIn: 'root',
})
export class CharactersService {

  public getCharacters(): Observable<Character[]> {
    // TODO: Mock, use real service
    const character: Character[] = [{
      abilities: [1, 2, 3, 4, 5, 6],
      ac: 10,
      conditions: [],
      hi: 0,
      hp: {
        val: 10,
        max: 10
      },
      id: 1,
      speed: 30,
    }]
    return of(character)
  }
}
