// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { ButtonDB } from './button-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class ButtonService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  ButtonServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private buttonsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.buttonsUrl = origin + '/api/gongmattree/go/v1/buttons';
  }

  /** GET buttons from the server */
  getButtons(): Observable<ButtonDB[]> {
    return this.http.get<ButtonDB[]>(this.buttonsUrl)
      .pipe(
        tap(_ => this.log('fetched buttons')),
        catchError(this.handleError<ButtonDB[]>('getButtons', []))
      );
  }

  /** GET button by id. Will 404 if id not found */
  getButton(id: number): Observable<ButtonDB> {
    const url = `${this.buttonsUrl}/${id}`;
    return this.http.get<ButtonDB>(url).pipe(
      tap(_ => this.log(`fetched button id=${id}`)),
      catchError(this.handleError<ButtonDB>(`getButton id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new button to the server */
  postButton(buttondb: ButtonDB): Observable<ButtonDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.post<ButtonDB>(this.buttonsUrl, buttondb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`posted buttondb id=${buttondb.ID}`)
      }),
      catchError(this.handleError<ButtonDB>('postButton'))
    );
  }

  /** DELETE: delete the buttondb from the server */
  deleteButton(buttondb: ButtonDB | number): Observable<ButtonDB> {
    const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
    const url = `${this.buttonsUrl}/${id}`;

    return this.http.delete<ButtonDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted buttondb id=${id}`)),
      catchError(this.handleError<ButtonDB>('deleteButton'))
    );
  }

  /** PUT: update the buttondb on the server */
  updateButton(buttondb: ButtonDB): Observable<ButtonDB> {
    const id = typeof buttondb === 'number' ? buttondb : buttondb.ID;
    const url = `${this.buttonsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    return this.http.put<ButtonDB>(url, buttondb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        this.log(`updated buttondb id=${buttondb.ID}`)
      }),
      catchError(this.handleError<ButtonDB>('updateButton'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
